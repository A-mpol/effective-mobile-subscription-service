package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	v1 "github.com/A-mpol/effective-mobile-subscription-service/internal/api/subscription/v1"
	"github.com/A-mpol/effective-mobile-subscription-service/internal/config"
	repoSubscription "github.com/A-mpol/effective-mobile-subscription-service/internal/repository/subscription"
	serviceSubscription "github.com/A-mpol/effective-mobile-subscription-service/internal/service/subscription"
	subscription_v1 "github.com/A-mpol/effective-mobile-subscription-service/pkg/proto/subscription/v1"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()
	cfg := config.Load(ctx)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GrpcServerConfig.Port()))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	defer func() {
		if cerr := lis.Close(); cerr != nil {
			log.Printf("failed to close listener: %v\n", cerr)
		}
	}()

	pool, err := pgxpool.New(ctx, cfg.PostgresConfig.URI())
	if err != nil {
		log.Fatalf("can not pgxpool.New: %v\n", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("failed to ping pgxpool: %v\n", err)
	}

	repo := repoSubscription.NewRepository(pool)
	service := serviceSubscription.NewService(repo)
	api := v1.NewApi(service)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			validator.UnaryServerInterceptor(),
			recovery.UnaryServerInterceptor(),
		),
	)

	subscription_v1.RegisterSubscriptionServiceServer(grpcServer, api)
	reflection.Register(grpcServer)

	go func() {
		log.Printf("gRPC server is listening on port: %d\n", cfg.GrpcServerConfig.Port())
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to grpcServer.Serve(): %v\n", err)
		}
	}()

	var gatewayServer *http.Server
	go func() {
		gatewayServer = startGrpcGatewayWithSwagger(ctx, cfg)
	}()

	_ = gatewayServer

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down gRPC server...")
	grpcServer.GracefulStop()
	log.Println("gRPC server stopped")
}

func startGrpcGatewayWithSwagger(ctx context.Context, cfg config.Config) *http.Server {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := subscription_v1.RegisterSubscriptionServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", cfg.GrpcServerConfig.Port()), opts); err != nil {
		log.Fatalf("failed to register gRPC-gateway: %v\n", err)
	}

	fileServer := http.FileServer(http.Dir("./api"))
	httpMux := http.NewServeMux()

	httpMux.Handle("/api/", mux)
	httpMux.Handle("/swagger-ui.html", fileServer)
	httpMux.Handle("/swagger.json", fileServer)

	httpMux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/swagger-ui.html", http.StatusMovedPermanently)
			return
		}
		fileServer.ServeHTTP(w, r)
	}))

	gatewayServer := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.HttpServerConfig.Port()),
		Handler:           httpMux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Printf("HTTP server with gRPC-Gateway and Swagger UI listening on port: %d\n", cfg.HttpServerConfig.Port())
	if err := gatewayServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to serve HTTP: %v\n", err)
	}

	return gatewayServer
}
