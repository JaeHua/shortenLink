# deploy/docker/show-rpc.Dockerfile
FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY . .
RUN go build -trimpath -ldflags="-s -w" -o /out/show-rpc cmd/show-rpc/show.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /out/show-rpc /app/show-rpc
COPY deploy/docker-compose/configs/show.yaml /app/etc/show.yaml
EXPOSE 8081
ENTRYPOINT ["/app/show-rpc","-f","/app/etc/show.yaml"]