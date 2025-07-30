package config

import (
	"context"
	"log"

	env_config "github.com/A-mpol/effective-mobile-subscription-service/internal/config/env"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type (
	// PostgresConfig -
	PostgresConfig interface {
		URI() string
	}

	// GrpcServerConfig -
	GrpcServerConfig interface {
		Port() int64
	}

	// HttpServerConfig -
	HttpServerConfig interface {
		Port() int64
	}

	// Config -
	Config struct {
		PostgresConfig   PostgresConfig
		GrpcServerConfig GrpcServerConfig
		HttpServerConfig HttpServerConfig
	}
)

func Load(ctx context.Context, path ...string) Config {
	if err := godotenv.Load(path...); err != nil {
		log.Println(ctx, "failed to godotenv.Load()", zap.Error(err))
	}

	postgresCfg, err := env_config.NewPostgresConfig()
	if err != nil {
		log.Fatal(ctx, "failed to init postgres config: %v", zap.Error(err))
	}

	grpcServerCfg, err := env_config.NewGrpcServerConfig()
	if err != nil {
		log.Fatal(ctx, "failed to init logger config: %v", zap.Error(err))
	}

	httpServerCfg, err := env_config.NewHttpServerConfig()
	if err != nil {
		log.Fatal(ctx, "failed to init http config: %v", zap.Error(err))
	}

	return Config{
		PostgresConfig:   postgresCfg,
		HttpServerConfig: httpServerCfg,
		GrpcServerConfig: grpcServerCfg,
	}
}
