package handler

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
	"shortenLink/cmd/show-api/internal/logic"
	"shortenLink/cmd/show-api/internal/svc"
	"shortenLink/cmd/show-api/internal/types"
)

// 根据短链 ShortUrl 跳转到长链
func ShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShowRequest
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
		l := logic.NewShowLogic(r.Context(), svcCtx)
		resp, err := l.Show(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			//httpx.OkJsonCtx(r.Context(), w, resp)
			http.Redirect(w, r, resp.LongUrl, http.StatusFound) // 302 跳转
		}
	}
}
