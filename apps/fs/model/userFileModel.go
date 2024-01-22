package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserFileModel = (*customUserFileModel)(nil)

type (
	// UserFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserFileModel.
	UserFileModel interface {
		userFileModel
		withSession(session sqlx.Session) UserFileModel
	}

	customUserFileModel struct {
		*defaultUserFileModel
	}
)

// NewUserFileModel returns a model for the database table.
func NewUserFileModel(conn sqlx.SqlConn) UserFileModel {
	return &customUserFileModel{
		defaultUserFileModel: newUserFileModel(conn),
	}
}

func (m *customUserFileModel) withSession(session sqlx.Session) UserFileModel {
	return NewUserFileModel(sqlx.NewSqlConnFromSession(session))
}
