package towerequipment

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTowerEquipmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTowerEquipmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTowerEquipmentLogic {
	return &DeleteTowerEquipmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTowerEquipmentLogic) DeleteTowerEquipment(req *types.DeleteTowerEquipmentReq) (resp *types.DeleteTowerEquipmentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
