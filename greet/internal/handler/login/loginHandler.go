package login

import (
	"net/http"

	"greet/common/respx"
	"greet/internal/logic/login"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := login.NewLoginLogic(r.Context(), svcCtx)
		resp, msg, err := l.Login(&req)
		respx.Response(w, resp, err, msg)
	}
}
