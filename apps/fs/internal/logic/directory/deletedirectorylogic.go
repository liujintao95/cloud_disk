package directory

import (
	"cloud_disk/apps/fs/internal/svc"
	"cloud_disk/apps/fs/internal/types"
	"cloud_disk/apps/fs/model"
	"cloud_disk/common/xerr"
	"context"
	"github.com/pkg/errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDirectoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDirectoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDirectoryLogic {
	return &DeleteDirectoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDirectoryLogic) DeleteDirectory(req *types.DeleteDirectoryReq) (resp *types.DeleteDirectoryResp, err error) {
	// 判断是否是强制，如果不是判断是否有子文件或子目录，有则报错，没有则删除
	// 如果强制递归删除子文件和目录
	var (
		dir          *model.UserDirectory
		dirChildren  []model.UserDirectory
		fileChildren []model.UserFileDetail
	)
	dir, err = l.svcCtx.UserDirectoryModel.FindOne(l.ctx, req.DirId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "数据库查询错误:%v, dir_id:%d", err, req.DirId)
	}
	if dir == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DIR_NOT_EXISTS), "当前文件夹不存在，dir_id:%d", req.DirId)
	}
	if !req.Force {
		dirChildren, err = l.svcCtx.UserDirectoryModel.FindAllByParentId(l.ctx, req.DirId)
		if err != nil && !errors.Is(err, model.ErrNotFound) {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "数据库查询子目录错误:%v, parent_id:%d", err, req.DirId)
		}
		if dirChildren != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DATA_LOCKED), "当前目录下存在子目录，无法进行删除:%v, dir_id:%d", err, req.DirId)
		}
		fileChildren, err = l.svcCtx.UserFileModel.FindAllByDirId(l.ctx, req.DirId)
		if err != nil && !errors.Is(err, model.ErrNotFound) {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "数据库查询目录文件错误:%v, dir_id:%d", err, req.DirId)
		}
		if fileChildren != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DATA_LOCKED), "当前目录下存在文件，无法进行删除:%v, dir_id:%d", err, req.DirId)
		}
	}
	dir.DeleteAt = time.Now()
	err = l.svcCtx.UserDirectoryModel.Update(l.ctx, dir)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "逻辑删除目录错误:%v, data:%#v", err, dir)
	}
	return
}
