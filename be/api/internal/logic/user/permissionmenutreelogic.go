package user

import (
	"context"

	"be/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionMenuTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPermissionMenuTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionMenuTreeLogic {
	return &PermissionMenuTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermissionMenuTreeLogic) PermissionMenuTree() error {
	// todo: add your logic here and delete this line

	return nil
}
