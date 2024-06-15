package user

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionPatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPermissionPatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionPatchLogic {
	return &PermissionPatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermissionPatchLogic) PermissionPatch(req *types.PatchPermissionReq) error {
	// todo: add your logic here and delete this line

	return nil
}
