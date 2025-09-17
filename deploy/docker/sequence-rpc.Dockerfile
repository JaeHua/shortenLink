# -------- Builder Stage --------
FROM golang:1.24-alpine AS builder

# 工作目录
WORKDIR /src

# 设置 Go 代理（优先 goproxy.cn，失败就直连）
ENV GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 先只复制 go.mod / go.sum，这样 go mod download 可以缓存依赖
COPY go.mod go.sum ./
RUN go mod download

# 再复制剩余代码
COPY . .

# 编译
RUN go build -trimpath -ldflags="-s -w" -o /out/sequence-rpc cmd/sequence-rpc/sequence.go

# -------- Runtime Stage --------
FROM alpine:3.20

WORKDIR /app

# 拷贝编译好的二进制
COPY --from=builder /out/sequence-rpc /app/sequence-rpc
COPY deploy/docker-compose/configs/sequence.yaml /app/etc/sequence.yaml

EXPOSE 8082

ENTRYPOINT ["/app/sequence-rpc", "-f", "/app/etc/sequence.yaml"]
