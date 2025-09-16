package logic

import (
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"shortenLink/cmd/show-rpc/internal/svc"
	"shortenLink/cmd/show-rpc/shortenLink/rpc/show"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	Err404 = errors.New("404")
)

type ShowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShowLogic) Show(in *show.ShowRequest) (*show.ShowResponse, error) {
	// 1. 先通过布隆过滤器判断
	exist, err := l.svcCtx.Filter.Exists([]byte(in.ShortCode))
	if err != nil {
		logx.Errorw("bloom filter exists failed", logx.Field("error", err), logx.Field("short_code", in.ShortCode))
		return nil, status.Error(codes.Aborted, "bloom filter exists failed")
	}
	if !exist {
		logx.Errorw("bloom filter not found", logx.Field("short_code", in.ShortCode))
		return nil, status.Error(codes.NotFound, "filter not found")
	}
	// 2. 去数据库查短链
	u, err := l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{
		String: in.ShortCode,
		Valid:  true,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logx.Errorw("shortCode not found in DB", logx.Field("short_code", in.ShortCode))
			return nil, status.Error(codes.NotFound, "shortCode not found in DB")
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl failed, shortCode", logx.Field("error", err), logx.Field("short_code", in.ShortCode))
		return nil, status.Error(codes.Aborted, "ShortUrlModel.FindOneBySurl failed")
	}
	return &show.ShowResponse{LongUrl: u.Lurl.String}, nil
}
