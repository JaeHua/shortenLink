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
RUN go build -trimpath -ldflags="-s -w" -o /out/show-api cmd/show-api/show.go

# -------- Runtime Stage --------
FROM alpine:3.20
WORKDIR /app

COPY --from=builder /out/show-api /app/show-api
COPY deploy/docker-compose/configs/show-api.yaml /app/etc/show-api.yaml

EXPOSE 8889
ENTRYPOINT ["/app/show-api","-f","/app/etc/show-api.yaml"]
