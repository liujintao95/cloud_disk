package file

import (
	"context"

	"cloud_disk/apps/fs/internal/svc"
	"cloud_disk/apps/fs/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDetailLogic {
	return &UserFileDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileDetailLogic) UserFileDetail(req *types.UserFileDetailReq) (resp *types.UserFileDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
