package directory

import (
	"cloud_disk/apps/fs/internal/svc"
	"cloud_disk/apps/fs/internal/types"
	"cloud_disk/apps/fs/model"
	"cloud_disk/common/constant"
	"cloud_disk/common/xerr"
	"context"
	"github.com/pkg/errors"

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
	var (
		currentDir *model.UserDirectory
		parentDir  *model.UserDirectory
	)
	currentDir, err = l.svcCtx.UserDirectoryModel.FindOne(l.ctx, req.DirId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "数据库查询错误:%v, dir_id:%d", err, req.DirId)
	}
	if currentDir == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DIR_NOT_EXISTS), "当前文件夹不存在，dir_id:%d", req.DirId)
	}
	if req.ParentId != constant.ROOT_DIR_ID {
		parentDir, err = l.svcCtx.UserDirectoryModel.FindOne(l.ctx, req.ParentId)
		if err != nil && !errors.Is(err, model.ErrNotFound) {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "数据库查询错误:%v, dir_id:%d", err, req.ParentId)
		}
		if parentDir == nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DIR_NOT_EXISTS), "目标文件夹不存在，dir_id:%d", req.ParentId)
		}
	}
	currentDir.ParentId = req.ParentId
	err = l.svcCtx.UserDirectoryModel.Update(l.ctx, currentDir)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "数据库更新错误:%v, data:%#v", err, currentDir)
	}
	return
}
