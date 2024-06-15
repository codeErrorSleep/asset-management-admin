package user

import (
	"context"

	"be/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCaptchaLogic {
	return &UserCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCaptchaLogic) UserCaptcha() error {
	// todo: add your logic here and delete this line

	return nil
}
