package file

import (
	"context"

	"cloud_disk/apps/fs/internal/svc"
	"cloud_disk/apps/fs/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MoveFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMoveFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoveFileLogic {
	return &MoveFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MoveFileLogic) MoveFile(req *types.MoveFileReq) (resp *types.MoveFileResp, err error) {
	// todo: add your logic here and delete this line

	return
}
