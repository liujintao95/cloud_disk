package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ UserFileModel = (*customUserFileModel)(nil)

type (
	// UserFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserFileModel.
	UserFileModel interface {
		userFileModel
		withSession(session sqlx.Session) UserFileModel
		FindAllByDirId(ctx context.Context, dirId int64) ([]UserFileDetail, error)
	}

	customUserFileModel struct {
		*defaultUserFileModel
	}

	UserFileDetail struct {
		Id       int64     `db:"id"`
		CreateAt time.Time `db:"create_at"`
		UpdateAt time.Time `db:"update_at"`
		DeleteAt time.Time `db:"delete_at"`
		UserId   int64     `db:"user_id"`
		FileId   int64     `db:"file_id"`
		DirId    int64     `db:"dir_id"`
		Name     string    `db:"name"`
		Ext      string    `db:"ext"`
		Size     int64     `db:"size"`
		Hash     string    `db:"hash"`
		Status   string    `db:"status"`
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

func (m *customUserFileModel) FindAllByDirId(ctx context.Context, dirId int64) ([]UserFileDetail, error) {
	sql := `
	select user_file.id as id, user_file.create_at as create_at, user_file.update_at as update_at, 
	       user_file.name as name, user_id, file_id, dir_id,  ext, "size", hash, status
	from user_file
	inner join file
	on user_file.file_id = file.id
	where parent_id = ?
	and user_file.delete_at is null 
	and file.delete_at is null 
	`
	var resp []UserFileDetail
	err := m.conn.QueryRowsCtx(ctx, &resp, sql, dirId)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
