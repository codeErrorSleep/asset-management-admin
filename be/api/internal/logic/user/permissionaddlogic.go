package user

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPermissionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionAddLogic {
	return &PermissionAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermissionAddLogic) PermissionAdd(req *types.AddPermissionReq) error {
	// todo: add your logic here and delete this line

	return nil
}
