# binary stage
FROM golang:1.14 AS builder

WORKDIR /app

ADD . .

WORKDIR /app/services/calc-api

RUN  CGO_ENABLED=0 go build -o /calc-api

# docker image  stage
FROM alpine:3.11

WORKDIR /app

COPY --from=builder /calc-api /app/

EXPOSE 3000
EXPOSE 10000

RUN GRPC_HEALTH_PROBE_VERSION=v0.3.1 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

ENTRYPOINT ./calc-api