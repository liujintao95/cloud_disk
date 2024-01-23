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
	// 判断是否是强制，如果不是判断是否有子文件或子目录，有则报错，没有则删除
	// 如果强制递归删除子文件和目录

	return
}
