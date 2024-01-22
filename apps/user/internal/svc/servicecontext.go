package svc

import (
	"cloud_disk/apps/user/internal/config"
	"cloud_disk/apps/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.UserModel
	sqlConn   sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlConn),
		sqlConn:   sqlConn,
	}
}
