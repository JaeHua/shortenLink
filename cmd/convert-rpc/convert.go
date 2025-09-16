package main

import (
	"flag"
	"fmt"
	"shortenLink/pkg/base62"

	"shortenLink/cmd/convert-rpc/internal/config"
	"shortenLink/cmd/convert-rpc/internal/server"
	"shortenLink/cmd/convert-rpc/internal/svc"
	"shortenLink/cmd/convert-rpc/shortenLink/rpc/convert"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/convert.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	base62.MustInit(c.BaseString)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		convert.RegisterConvertServer(grpcServer, server.NewConvertServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
