package file

import (
	"net/http"

	"cloud_disk/apps/fs/internal/logic/file"
	"cloud_disk/apps/fs/internal/svc"
	"cloud_disk/apps/fs/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MoveFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MoveFileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewMoveFileLogic(r.Context(), svcCtx)
		resp, err := l.MoveFile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
