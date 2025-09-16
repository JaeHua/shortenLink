# deploy/docker/convert-rpc.Dockerfile
FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY . .
RUN go build -trimpath -ldflags="-s -w" -o /out/convert-rpc cmd/convert-rpc/convert.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /out/convert-rpc /app/convert-rpc
COPY deploy/docker-compose/configs/convert.yaml /app/etc/convert.yaml
EXPOSE 8080
ENTRYPOINT ["/app/convert-rpc","-f","/app/etc/convert.yaml"]