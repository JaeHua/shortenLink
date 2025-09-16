# deploy/docker/sequence-rpc.Dockerfile
FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY . .
RUN go build -trimpath -ldflags="-s -w" -o /out/sequence-rpc cmd/sequence-rpc/sequence.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /out/sequence-rpc /app/sequence-rpc
COPY deploy/docker-compose/configs/sequence.yaml /app/etc/sequence.yaml
EXPOSE 8082
ENTRYPOINT ["/app/sequence-rpc","-f","/app/etc/sequence.yaml"]