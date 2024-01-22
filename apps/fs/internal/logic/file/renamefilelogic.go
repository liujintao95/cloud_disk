package file

import (
	"context"

	"cloud_disk/apps/fs/internal/svc"
	"cloud_disk/apps/fs/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RenameFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRenameFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RenameFileLogic {
	return &RenameFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RenameFileLogic) RenameFile(req *types.RenameFileReq) (resp *types.RenameFileResp, err error) {
	// todo: add your logic here and delete this line

	return
}
