package user

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleRemoveUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleRemoveUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleRemoveUserLogic {
	return &RoleRemoveUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleRemoveUserLogic) RoleRemoveUser(req *types.PatchRoleOperateUserReq) error {
	// todo: add your logic here and delete this line

	return nil
}
