package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"shortenLink/cmd/convert-api/internal/logic"
	"shortenLink/cmd/convert-api/internal/svc"
	"shortenLink/cmd/convert-api/internal/types"
)

func ConvertHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ConvertRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// 参数规则校验
		if err := validator.New().StructCtx(r.Context(), &req); err != nil {
			logx.Errorw("参数校验失败", logx.Field("error", err.Error()), logx.Field("req", req))
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewConvertLogic(r.Context(), svcCtx)
		resp, err := l.Convert(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
