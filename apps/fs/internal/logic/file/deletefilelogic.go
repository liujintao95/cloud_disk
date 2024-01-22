package file

import (
	"context"

	"cloud_disk/apps/fs/internal/svc"
	"cloud_disk/apps/fs/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFileLogic {
	return &DeleteFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFileLogic) DeleteFile(req *types.DeleteFileReq) (resp *types.DeleteFileResp, err error) {
	// todo: add your logic here and delete this line

	return
}
