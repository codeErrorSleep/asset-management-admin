package equipmentclass

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListEquipmentClassesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListEquipmentClassesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListEquipmentClassesLogic {
	return &ListEquipmentClassesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListEquipmentClassesLogic) ListEquipmentClasses(req *types.ListEquipmentClassesReq) (resp *types.ListEquipmentClassesResp, err error) {
	// todo: add your logic here and delete this line

	return
}
