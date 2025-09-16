package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"shortenLink/cmd/convert-api/internal/errorx"
	"shortenLink/cmd/convert-rpc/shortenLink/rpc/convert"

	"shortenLink/cmd/convert-api/internal/svc"
	"shortenLink/cmd/convert-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// 调用 convert-rpc
	rpcResp, err := l.svcCtx.ConvertRpc.Convert(l.ctx, &convert.ConvertRequest{
		LongUrl: req.LongUrl,
	})
	if err != nil {
		logx.Errorw("call convert-rpc failed", logx.Field("err", err))
		if st, ok := status.FromError(err); ok {
			// 返回自定义 CodeError
			return nil, errorx.NewCodeError(errorx.RPCErrCode, st.Message())
		}
		return nil, errorx.NewCodeError(errorx.InternalErrCode, err.Error())
	}

	// 返回给 API 层
	return &types.ConvertResponse{
		ShortUrl: rpcResp.ShortUrl,
	}, nil
	return
}
