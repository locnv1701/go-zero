package user

import (
	"context"
	"errors"

	"greet/internal/svc"
	"greet/internal/types"
	"greet/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoRes, err error) {
	// todo: add your logic here and delete this line
	DB := l.svcCtx.DB

	userRepo := models.NewUserModel(DB)

	user, err := userRepo.FindOne(l.ctx, int64(req.Uid))

	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return nil, errors.New("User not exist")
		}
		return nil, errors.New("FindOne")
	}

	userInfo := &types.User{
		Username: user.Name,
		Uid:      uint(user.Id),
	}

	return &types.UserInfoRes{
		UserInfo: *userInfo,
	}, err
}
