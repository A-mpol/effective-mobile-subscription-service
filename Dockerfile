FROM golang:1.24.2-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY .env .
RUN CGO_ENABLED=0 GOOS=linux go build -o /subscription-service ./cmd

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /subscription-service .
EXPOSE 8081
CMD ["./subscription-service"]
