package equipmentclass

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEquipmentClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateEquipmentClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEquipmentClassLogic {
	return &UpdateEquipmentClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateEquipmentClassLogic) UpdateEquipmentClass(req *types.UpdateEquipmentClassReq) (resp *types.UpdateEquipmentClassResp, err error) {
	// todo: add your logic here and delete this line

	return
}
