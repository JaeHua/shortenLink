package svc

import (
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"shortenLink/cmd/convert-rpc/internal/config"
	"shortenLink/cmd/sequence-rpc/shortenLink/rpc/sequence"
	"shortenLink/model"
)

type ServiceContext struct {
	Config            config.Config
	ShortUrlModel     model.ShortUrlMapModel
	ShortUrlBlacklist map[string]struct{}
	SequenceRpc       sequence.SequenceClient
	Filter            *bloom.Filter
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	rpcConn := zrpc.MustNewClient(c.SequenceRpc)
	// 将黑名单列表转换为map，方便后续查询
	m := make(map[string]struct{}, len(c.ShortUrlBlacklist))
	for _, v := range c.ShortUrlBlacklist {
		m[v] = struct{}{} // 不占用额外内存
	}
	// 初始化布隆过滤器
	store := redis.New(c.CacheRedis[0].Host, func(r *redis.Redis) {
		r.Type = redis.NodeType
	})
	filter := bloom.New(store, "short_url_bloom_filter", 20*(1<<20))
	return &ServiceContext{
		Config:            c,
		ShortUrlModel:     model.NewShortUrlMapModel(conn, c.CacheRedis),
		ShortUrlBlacklist: m,
		SequenceRpc:       sequence.NewSequenceClient(rpcConn.Conn()),
		Filter:            filter,
	}
}
