package logic

import (
	"context"

	"shortenLink/cmd/convert-rpc/internal/svc"
	"shortenLink/cmd/convert-rpc/shortenLink/rpc/convert"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConvertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConvertLogic) Convert(in *convert.ConvertRequest) (*convert.ConvertResponse, error) {
	// 1. 校验输入的数据
	// 1.1 数据不能空
	// validator做参数校验
	//上一层已经解决，不会执行逻辑
	// 1.2 数据必须是合法的url

	return &convert.ConvertResponse{}, nil
}
