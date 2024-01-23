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
	var (
		dirResult  []model.UserDirectory
		fileResult []model.UserFileDetail
		fileList   []types.File
		dirList    []types.Directory
	)
	dirResult, err = l.svcCtx.UserDirectoryModel.FindAllByParentId(l.ctx, req.ParentId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "数据库查询子目录错误:%v, parent_id:%d", err, req.ParentId)
	}
	fileResult, err = l.svcCtx.UserFileModel.FindAllByDirId(l.ctx, req.ParentId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "数据库查询目录文件错误:%v, dir_id:%d", err, req.ParentId)
	}
	if dirResult != nil {
		for _, dir := range dirResult {
			dirList = append(dirList, types.Directory{
				Id:       dir.Id,
				Name:     dir.Name,
				ParentId: dir.ParentId,
			})
		}
	}
	if fileResult != nil {
		for _, file := range fileResult {
			fileList = append(fileList, types.File{
				Id:     file.Id,
				Name:   file.Name,
				Size:   file.Size,
				Ext:    file.Ext,
				Hash:   file.Hash,
				Status: file.Status,
			})
		}
	}
	resp.FileList = fileList
	resp.DirectoryList = dirList
	return
}
