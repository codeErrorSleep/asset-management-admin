package user

import (
	"context"

	"be/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdatePasswordLogic {
	return &UserUpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdatePasswordLogic) UserUpdatePassword() error {
	// todo: add your logic here and delete this line

	return nil
}
