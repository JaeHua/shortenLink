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
RUN go build -trimpath -ldflags="-s -w" -o /out/convert-rpc cmd/convert-rpc/convert.go

# -------- Runtime Stage --------
FROM alpine:3.20
WORKDIR /app
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
COPY --from=builder /out/convert-rpc /app/convert-rpc
COPY deploy/docker-compose/configs/convert.yaml /app/etc/convert.yaml

EXPOSE 8080
ENTRYPOINT ["/app/convert-rpc","-f","/app/etc/convert.yaml"]
