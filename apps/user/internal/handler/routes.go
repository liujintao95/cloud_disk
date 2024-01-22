// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "cloud_disk/apps/user/internal/handler/user"
	"cloud_disk/apps/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/detail",
				Handler: user.DetailHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1/user"),
	)
}