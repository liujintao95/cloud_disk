package directory

import (
	"net/http"

	"cloud_disk/apps/fs/internal/logic/directory"
	"cloud_disk/apps/fs/internal/svc"
	"cloud_disk/apps/fs/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateDirectoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateDirectoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := directory.NewCreateDirectoryLogic(r.Context(), svcCtx)
		resp, err := l.CreateDirectory(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
