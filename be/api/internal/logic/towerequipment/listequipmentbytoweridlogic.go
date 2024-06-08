package towerequipment

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListEquipmentByTowerIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListEquipmentByTowerIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListEquipmentByTowerIdLogic {
	return &ListEquipmentByTowerIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListEquipmentByTowerIdLogic) ListEquipmentByTowerId(req *types.ListEquipmentByTowerIdReq) (resp *types.ListEquipmentByTowerIdResp, err error) {
	// todo: add your logic here and delete this line

	return
}
