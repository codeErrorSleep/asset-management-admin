package user

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddUserLogic {
	return &RoleAddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleAddUserLogic) RoleAddUser(req *types.PatchRoleOperateUserReq) error {
	// todo: add your logic here and delete this line

	return nil
}
