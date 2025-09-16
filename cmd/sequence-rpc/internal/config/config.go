package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	SequenceDB SequenceDB
	CacheRedis cache.CacheConf
}
type SequenceDB struct {
	DSN string
}
