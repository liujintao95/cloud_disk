package directory

import (
	"net/http"

	"cloud_disk/apps/fs/internal/logic/directory"
	"cloud_disk/apps/fs/internal/svc"
	"cloud_disk/apps/fs/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RemaneDirectoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemaneDirectoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := directory.NewRemaneDirectoryLogic(r.Context(), svcCtx)
		resp, err := l.RemaneDirectory(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
