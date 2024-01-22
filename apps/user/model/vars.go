package model

import (
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	ErrNotFound        = sqlx.ErrNotFound
	ErrInvalidObjectId = errors.New("invalid objectId")
)
