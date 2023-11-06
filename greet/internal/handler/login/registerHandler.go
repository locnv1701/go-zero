package login

import (
	"net/http"

	"greet/common/respx"
	"greet/internal/logic/login"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := login.NewRegisterLogic(r.Context(), svcCtx)
		msg, err := l.Register(&req)
		respx.Response(w, "", err, msg)
	}
}
