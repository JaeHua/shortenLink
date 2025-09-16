package logic

import (
	"context"

	"shortenLink/cmd/show-rpc/internal/svc"
	"shortenLink/cmd/show-rpc/shortenLink/rpc/show"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return &show.ShowResponse{}, nil
}
