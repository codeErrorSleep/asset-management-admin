package user

import (
	"context"

	"be/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListPageLogic {
	return &RoleListPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListPageLogic) RoleListPage() error {
	// todo: add your logic here and delete this line

	return nil
}
