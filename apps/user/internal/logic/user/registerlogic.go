package user

import (
	"cloud_disk/apps/user/model"
	"cloud_disk/common/ctxdata"
	"cloud_disk/common/tool"
	"cloud_disk/common/xerr"
	"context"
	"database/sql"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"time"

	"cloud_disk/apps/user/internal/svc"
	"cloud_disk/apps/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	var (
		user         *model.User
		userId       int64
		insertResult sql.Result
		now          int64
	)
	user, err = l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "数据库查询错误:%v, 手机号:%+v", err, req.Mobile)
	}
	if user != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_ALREADY_EXISTS), "该手机号已被注册:%s,err:%v", req.Mobile, err)
	}
	user = new(model.User)
	user.Mobile = req.Mobile
	user.Nickname = req.Nickname
	user.Password = tool.Md5ByString(req.Password)
	insertResult, err = l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "数据库写入错误:%v, 用户:%+v", err, user)
	}
	userId, err = insertResult.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "获取自增ID错误:%v, 用户:%+v", err, user)
	}
	resp.AccessToken, err = l.GenerateToken(userId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR), "生成userId : %d", userId)
	}
	now = time.Now().Unix()
	resp.AccessExpire = now + l.svcCtx.Config.JwtAuth.AccessExpire
	resp.RefreshAfter = now + l.svcCtx.Config.JwtAuth.AccessExpire/2
	return
}

func (l *RegisterLogic) GenerateToken(userId int64) (string, error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, userId)
	if err != nil {
		return "", errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR), "生成JwtToken错误:%v", err)
	}

	return accessToken, nil
}

func (l *RegisterLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[ctxdata.CtxKeyJwtUserId] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
