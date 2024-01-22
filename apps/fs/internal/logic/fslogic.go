package logic

import (
	"cloud_disk/apps/fs/internal/svc"
	"cloud_disk/apps/fs/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FsLogic {
	return &FsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FsLogic) Fs(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
