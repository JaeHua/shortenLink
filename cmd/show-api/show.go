package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"shortenLink/cmd/show-api/internal/errorx"

	"shortenLink/cmd/show-api/internal/config"
	"shortenLink/cmd/show-api/internal/handler"
	"shortenLink/cmd/show-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/show-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	// 注册自定义的错误处理器
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, any) {
		var e *errorx.CodeError
		switch {
		case errors.As(err, &e):
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
