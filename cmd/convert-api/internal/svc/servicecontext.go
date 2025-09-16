package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"shortenLink/cmd/convert-api/internal/config"
	"shortenLink/cmd/convert-rpc/convertclient"
)

type ServiceContext struct {
	Config     config.Config
	ConvertRpc convertclient.Convert
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ConvertRpc: convertclient.NewConvert(zrpc.MustNewClient(c.ConvertRpc)),
	}
}
