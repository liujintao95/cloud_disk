// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	directory "cloud_disk/apps/fs/internal/handler/directory"
	file "cloud_disk/apps/fs/internal/handler/file"
	"cloud_disk/apps/fs/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create_directory",
				Handler: directory.CreateDirectoryHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete_directory",
				Handler: directory.DeleteDirectoryHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/move_directory",
				Handler: directory.MoveDirectoryHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/rename_directory",
				Handler: directory.RemaneDirectoryHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user_directory_list",
				Handler: directory.UserDirectoryListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1/file"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/delete_file",
				Handler: file.DeleteFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/move_file",
				Handler: file.MoveFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/rename_file",
				Handler: file.RenameFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user_file_detail",
				Handler: file.UserFileDetailHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1/file"),
	)
}