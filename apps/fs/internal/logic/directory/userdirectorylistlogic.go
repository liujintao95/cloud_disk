package directory

import (
	"context"

	"cloud_disk/apps/fs/internal/svc"
	"cloud_disk/apps/fs/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDirectoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDirectoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDirectoryListLogic {
	return &UserDirectoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDirectoryListLogic) UserDirectoryList(req *types.UserDirectoryListReq) (resp *types.UserDirectoryListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
