package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shortenLink/cmd/sequence-rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Conn   sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Conn:   sqlx.NewMysql(c.SequenceDB.DSN),
	}
}
