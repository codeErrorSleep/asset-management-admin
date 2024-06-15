package user

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPasswordLogic {
	return &UserPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPasswordLogic) UserPassword(req *types.AuthPwReq) error {
	// todo: add your logic here and delete this line

	return nil
}
