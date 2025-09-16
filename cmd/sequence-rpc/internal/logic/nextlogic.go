package logic

import (
	"context"

	"shortenLink/cmd/sequence-rpc/internal/svc"
	"shortenLink/cmd/sequence-rpc/shortenLink/rpc/sequence"

	"github.com/zeromicro/go-zero/core/logx"
)

type NextLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNextLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NextLogic {
	return &NextLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NextLogic) Next(in *sequence.Empty) (*sequence.NextResp, error) {
	// todo: add your logic here and delete this line

	return &sequence.NextResp{}, nil
}
