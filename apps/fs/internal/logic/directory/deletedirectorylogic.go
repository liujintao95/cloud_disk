package directory

import (
	"context"

	"cloud_disk/apps/fs/internal/svc"
	"cloud_disk/apps/fs/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDirectoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDirectoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDirectoryLogic {
	return &DeleteDirectoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDirectoryLogic) DeleteDirectory(req *types.DeleteDirectoryReq) (resp *types.DeleteDirectoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
