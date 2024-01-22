package svc

import (
	"cloud_disk/apps/fs/internal/config"
	"cloud_disk/apps/fs/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	FileModel          model.FileModel
	UserFileModel      model.UserFileModel
	UserDirectoryModel model.UserDirectoryModel
	sqlConn            sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:             c,
		FileModel:          model.NewFileModel(sqlConn),
		UserFileModel:      model.NewUserFileModel(sqlConn),
		UserDirectoryModel: model.NewUserDirectoryModel(sqlConn),
		sqlConn:            sqlConn,
	}
}
