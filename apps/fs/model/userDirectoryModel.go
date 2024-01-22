package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserDirectoryModel = (*customUserDirectoryModel)(nil)

type (
	// UserDirectoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserDirectoryModel.
	UserDirectoryModel interface {
		userDirectoryModel
		withSession(session sqlx.Session) UserDirectoryModel
	}

	customUserDirectoryModel struct {
		*defaultUserDirectoryModel
	}
)

// NewUserDirectoryModel returns a model for the database table.
func NewUserDirectoryModel(conn sqlx.SqlConn) UserDirectoryModel {
	return &customUserDirectoryModel{
		defaultUserDirectoryModel: newUserDirectoryModel(conn),
	}
}

func (m *customUserDirectoryModel) withSession(session sqlx.Session) UserDirectoryModel {
	return NewUserDirectoryModel(sqlx.NewSqlConnFromSession(session))
}
