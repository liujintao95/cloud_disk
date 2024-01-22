package directory

import (
	"context"

	"cloud_disk/apps/fs/internal/svc"
	"cloud_disk/apps/fs/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDirectoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDirectoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDirectoryLogic {
	return &CreateDirectoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDirectoryLogic) CreateDirectory(req *types.CreateDirectoryReq) (resp *types.CreateDirectoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
