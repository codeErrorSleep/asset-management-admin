package user

import (
	"context"

	"be/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionListTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPermissionListTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionListTreeLogic {
	return &PermissionListTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermissionListTreeLogic) PermissionListTree() error {
	// todo: add your logic here and delete this line

	return nil
}
