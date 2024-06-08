package towerequipment

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTowerEquipmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTowerEquipmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTowerEquipmentLogic {
	return &CreateTowerEquipmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTowerEquipmentLogic) CreateTowerEquipment(req *types.CreateTowerEquipmentReq) (resp *types.CreateTowerEquipmentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
