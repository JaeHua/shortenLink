package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	ShortUrlDB  ShortUrlDB
	SequenceRpc zrpc.RpcClientConf
	BaseString  string

	ShortUrlBlacklist []string

	ShortDomain string

	CacheRedis cache.CacheConf
}
type ShortUrlDB struct {
	DSN string
}
