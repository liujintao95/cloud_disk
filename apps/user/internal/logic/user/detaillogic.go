package user

import (
	"cloud_disk/apps/user/model"
	"cloud_disk/common/ctxdata"
	"cloud_disk/common/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"cloud_disk/apps/user/internal/svc"
	"cloud_disk/apps/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	var (
		userId   int64
		user     *model.User
		respUser types.User
	)
	userId = ctxdata.GetUidFromCtx(l.ctx)
	user, err = l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "获取登录用户信息失败:%v, id:%d", err, userId)
	}
	if user == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_EXISTS), "用户不存在，ID:%d", userId)
	}
	_ = copier.Copy(&respUser, user)
	resp.UserInfo = respUser
	return
}
