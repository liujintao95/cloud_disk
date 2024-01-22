package directory

import (
	"context"

	"cloud_disk/apps/fs/internal/svc"
	"cloud_disk/apps/fs/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemaneDirectoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemaneDirectoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemaneDirectoryLogic {
	return &RemaneDirectoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemaneDirectoryLogic) RemaneDirectory(req *types.RemaneDirectoryReq) (resp *types.RemaneDirectoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
