package user

import (
	"context"

	"be/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type RolePermissionsTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRolePermissionsTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RolePermissionsTreeLogic {
	return &RolePermissionsTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RolePermissionsTreeLogic) RolePermissionsTree() error {
	// todo: add your logic here and delete this line

	return nil
}
