# -------- Builder Stage --------
FROM golang:1.24-alpine AS builder
WORKDIR /src

ENV GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags="-s -w" -o /out/convert-api cmd/convert-api/convert.go

# -------- Runtime Stage --------
FROM alpine:3.20
WORKDIR /app

COPY --from=builder /out/convert-api /app/convert-api
COPY deploy/docker-compose/configs/convert-api.yaml /app/etc/convert-api.yaml

EXPOSE 8888
ENTRYPOINT ["/app/convert-api","-f","/app/etc/convert-api.yaml"]
