// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"greet/internal/svc"
	"greet/internal/handler/login"

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
}
