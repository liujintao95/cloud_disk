package user

import (
	"cloud_disk/apps/user/model"
	"cloud_disk/common/tool"
	"cloud_disk/common/xerr"
	"context"
	"github.com/pkg/errors"
	"time"

	"cloud_disk/apps/user/internal/svc"
	"cloud_disk/apps/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	var (
		user          *model.User
		registerLogic *RegisterLogic
		now           int64
	)
	user, err = l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据手机号查询用户信息失败:%v, 手机号:%s", err, req.Mobile)
	}
	if user == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_MOBILE_PWD_ERROR), "手机号:%s", req.Mobile)
	}
	if !(tool.Md5ByString(req.Password) == user.Password) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_MOBILE_PWD_ERROR), "手机号:%s", req.Mobile)
	}
	registerLogic = NewRegisterLogic(l.ctx, l.svcCtx)
	resp.AccessToken, err = registerLogic.GenerateToken(user.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR), "生成userId : %d", user.Id)
	}
	now = time.Now().Unix()
	resp.AccessExpire = now + l.svcCtx.Config.JwtAuth.AccessExpire
	resp.RefreshAfter = now + l.svcCtx.Config.JwtAuth.AccessExpire/2
	return
}
