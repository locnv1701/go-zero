package login

import (
	"context"
	"errors"

	"greet/common/helper"
	"greet/common/respx"
	"greet/internal/svc"
	"greet/internal/types"
	"greet/models"

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

func (l *RegisterLogic) Register(req *types.RegisterReq) (msg respx.SucMsg, err error) {
	err = l.register(req)
	if err != nil {
		return msg, err
	}

	return respx.SucMsg{Msg: "Succesfully"}, nil
}

func (l *RegisterLogic) register(req *types.RegisterReq) (err error) {
	// todo: add your logic here and delete this line

	DB := l.svcCtx.DB

	userRepo := models.NewUserModel(DB)

	userRegister := models.User{
		Name:     req.Username,
		Password: helper.MakeHash(req.Password),
		Tel:      int64(req.Tel),
	}

	_, err = userRepo.FindOneByName(l.ctx, req.Username)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			_, err := userRepo.Insert(l.ctx, &userRegister)
			if err != nil {
				return err
			}
			return nil
		}
		return errors.New("FindOneByName")
	}

	return errors.New("username exist")
}
