package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"shortenLink/cmd/show-api/internal/config"
	"shortenLink/cmd/show-rpc/showclient"
)

type ServiceContext struct {
	Config  config.Config
	ShowRpc showclient.Show
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		ShowRpc: showclient.NewShow(zrpc.MustNewClient(c.ShowRpc)),
	}
}
