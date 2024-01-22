package handler

import (
	"cloud_disk/apps/fs/internal/logic"
	"cloud_disk/apps/fs/internal/svc"
	"cloud_disk/apps/fs/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFsLogic(r.Context(), svcCtx)
		resp, err := l.Fs(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
