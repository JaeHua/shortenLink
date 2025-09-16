# deploy/docker/convert-api.Dockerfile
FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY . .
RUN go build -trimpath -ldflags="-s -w" -o /out/convert-api cmd/convert-api/convert.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /out/convert-api /app/convert-api
COPY deploy/docker-compose/configs/convert-api.yaml /app/etc/convert-api.yaml
EXPOSE 8888
ENTRYPOINT ["/app/convert-api","-f","/app/etc/convert-api.yaml"]