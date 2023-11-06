// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"greet/internal/handler/ws"
	"greet/internal/handler/login"
	"greet/internal/handler/user"
	"greet/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/greet/from/:name",
				Handler: GreetHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,		
				Path:    "/login",
				Handler: login.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,		
				Path:    "/register",
				Handler: login.RegisterHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/info",
				Handler: user.UserInfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/users"),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ws",
				Handler: ws.WsHandler(serverCtx),
			},
		},
	)

}
