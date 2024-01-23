package directory

import (
	"cloud_disk/apps/fs/model"
	"cloud_disk/common/xerr"
	"context"
	"github.com/pkg/errors"

	"cloud_disk/apps/fs/internal/svc"
	"cloud_disk/apps/fs/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RenameDirectoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRenameDirectoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RenameDirectoryLogic {
	return &RenameDirectoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RenameDirectoryLogic) RenameDirectory(req *types.RenameDirectoryReq) (resp *types.RenameDirectoryResp, err error) {
	var (
		dir *model.UserDirectory
	)
	dir, err = l.svcCtx.UserDirectoryModel.FindOne(l.ctx, req.DirId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "数据库查询错误:%v, dir_id:%d", err, req.DirId)
	}
	if dir == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DIR_NOT_EXISTS), "当前文件夹不存在，dir_id:%d", req.DirId)
	}
	dir.Name = req.Rename
	err = l.svcCtx.UserDirectoryModel.Update(l.ctx, dir)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "数据库更新错误:%v, data:%#v", err, dir)
	}
	return
}
