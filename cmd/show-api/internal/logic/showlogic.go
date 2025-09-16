package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"shortenLink/cmd/show-api/internal/errorx"
	"shortenLink/cmd/show-rpc/shortenLink/rpc/show"

	"shortenLink/cmd/show-api/internal/svc"
	"shortenLink/cmd/show-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据短链 ShortUrl 跳转到长链
func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// 调用 show-rpc
	rpcResp, err := l.svcCtx.ShowRpc.Show(l.ctx, &show.ShowRequest{
		ShortCode: req.ShortUrl, // 注意字段对应
	})
	if err != nil {
		logx.Errorw("call show-rpc failed, shortUrl", logx.Field("shortUrl", req.ShortUrl), logx.Field("error", err.Error()))
		// 解析 gRPC 错误
		if st, ok := status.FromError(err); ok {
			return nil, errorx.NewCodeError(errorx.RPCErrCode, st.Message())
		}

		// 其他错误返回默认 CodeError
		return nil, errorx.NewCodeError(errorx.InternalErrCode, err.Error())
	}

	// 返回给 API 层
	return &types.ShowResponse{
		LongUrl: rpcResp.LongUrl,
	}, nil
	return
}
