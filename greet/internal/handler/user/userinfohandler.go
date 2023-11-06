package user

import (
	"net/http"
	"strings"

	"greet/common/helper"
	"greet/internal/logic/user"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := strings.Trim(strings.Split(r.Header.Get("Authorization"), "Bearer")[1], " ")
		payload, err := helper.ParseJwtToken(token, svcCtx.Config.Auth.AccessSecret)

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}

		userInfoReq := types.UserInfoReq{
			Username: payload.Username,
			Uid:      uint32(payload.Uid),
		}

		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&userInfoReq)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
