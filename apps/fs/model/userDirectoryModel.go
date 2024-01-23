package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserDirectoryModel = (*customUserDirectoryModel)(nil)

type (
	// UserDirectoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserDirectoryModel.
	UserDirectoryModel interface {
		userDirectoryModel
		withSession(session sqlx.Session) UserDirectoryModel
		FindOneByNameParentId(ctx context.Context, name string, parentId int64) (*UserDirectory, error)
		FindAllByParentId(ctx context.Context, parentId int64) ([]UserDirectory, error)
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

func (m *customUserDirectoryModel) FindOneByNameParentId(ctx context.Context, name string, parentId int64) (*UserDirectory, error) {
	query := fmt.Sprintf("select %s from %s where `name` = ? and `parent_id` = ? and delete_at is null limit 1", userDirectoryRows, m.table)
	var resp UserDirectory
	err := m.conn.QueryRowCtx(ctx, &resp, query, name, parentId)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUserDirectoryModel) FindAllByParentId(ctx context.Context, parentId int64) ([]UserDirectory, error) {
	query := fmt.Sprintf("select %s from %s where `parent_id` = ? and delete_at is null", userDirectoryRows, m.table)
	var resp []UserDirectory
	err := m.conn.QueryRowCtx(ctx, &resp, query, parentId)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
