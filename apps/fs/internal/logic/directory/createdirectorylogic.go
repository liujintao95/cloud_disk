package directory

import (
	"cloud_disk/apps/fs/model"
	"cloud_disk/common/ctxdata"
	"cloud_disk/common/xerr"
	"context"
	"github.com/pkg/errors"

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
	var (
		dir *model.UserDirectory
	)
	dir, err = l.svcCtx.UserDirectoryModel.FindOneByNameParentId(l.ctx, req.Name, req.ParentId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "数据库查询错误:%v, query:%#v", err, req)
	}
	if dir != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_ALREADY_EXISTS), "当前目录下存在同名文件夹:%s", req.Name)
	}
	dir = new(model.UserDirectory)
	dir.Name = req.Name
	dir.ParentId = req.ParentId
	dir.UserId = ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.UserDirectoryModel.Insert(l.ctx, dir)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "数据库写入错误:%v, data:%#v", err, dir)
	}
	return
}
