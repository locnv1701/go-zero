package login

import (
	"context"
	"errors"
	"greet/common/helper"
	"greet/common/respx"
	"greet/internal/svc"
	"greet/internal/token"
	"greet/internal/types"
	"greet/models"

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

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, msg respx.SucMsg, err error) {

	userModel := models.NewUserModel(l.svcCtx.DB)

	user, err := userModel.FindOneByNameAndPassword(l.ctx, req.Username, helper.MakeHash(req.Password))

	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return nil, respx.SucMsg{
				Msg: "Username or password is incorrect",
			}, errors.New("username or password is incorrect")
		} else {
			return nil, respx.SucMsg{
				Msg: "FindOneByNameAndPassword",
			}, errors.New("FindOneByNameAndPassword")
		}
	}

	jwt, _ := helper.GenerateJwtToken(
		&helper.GenerateJwtStruct{
			Uid:      uint(user.Id),
			Username: user.Name,
		},
		l.svcCtx.Config.Auth.AccessSecret,
		l.svcCtx.Config.Auth.AccessExpire,
	)

	token.AddTokenToRedis(int(user.Id), jwt)

	return &types.LoginRes{
		UserId: int(user.Id),
		Token:  jwt,
	}, respx.SucMsg{Msg: "Successfully"}, nil

}
