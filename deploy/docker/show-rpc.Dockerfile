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
RUN go build -trimpath -ldflags="-s -w" -o /out/show-rpc cmd/show-rpc/show.go

# -------- Runtime Stage --------
FROM alpine:3.20
WORKDIR /app

COPY --from=builder /out/show-rpc /app/show-rpc
COPY deploy/docker-compose/configs/show.yaml /app/etc/show.yaml

EXPOSE 8081
ENTRYPOINT ["/app/show-rpc","-f","/app/etc/show.yaml"]
