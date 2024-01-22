package directory

import (
	"context"

	"cloud_disk/apps/fs/internal/svc"
	"cloud_disk/apps/fs/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MoveDirectoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMoveDirectoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoveDirectoryLogic {
	return &MoveDirectoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MoveDirectoryLogic) MoveDirectory(req *types.MoveDirectoryReq) (resp *types.MoveDirectoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
