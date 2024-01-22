package model

import (
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/net/context"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		withSession(session sqlx.Session) UserModel
		TransInsert(ctx context.Context, data *User, session sqlx.Session) (sql.Result, error)
		TransUpdate(ctx context.Context, newData *User, session sqlx.Session) error
		TransDelete(ctx context.Context, id int64, session sqlx.Session) error
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

func (m *customUserModel) withSession(session sqlx.Session) UserModel {
	return NewUserModel(sqlx.NewSqlConnFromSession(session))
}

func (m *defaultUserModel) TransInsert(ctx context.Context, data *User, session sqlx.Session) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
	ret, err := session.ExecCtx(ctx, query, data.DeleteAt, data.Mobile, data.Password, data.Nickname, data.Avatar, data.Info)
	return ret, err
}

func (m *defaultUserModel) TransUpdate(ctx context.Context, newData *User, session sqlx.Session) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
	_, err := session.ExecCtx(ctx, query, newData.DeleteAt, newData.Mobile, newData.Password, newData.Nickname, newData.Avatar, newData.Info, newData.Id)
	return err
}

func (m *defaultUserModel) TransDelete(ctx context.Context, id int64, session sqlx.Session) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := session.ExecCtx(ctx, query, id)
	return err
}
