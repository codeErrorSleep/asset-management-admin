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
	respEquipmentList := make([]types.EquipmentDetail, 0)
	resp = &types.ListEquipmentByTowerIdResp{
		PageData: respEquipmentList,
	}

	// 直接查关联表,
	towerEquipmentRelateIDList, err := l.svcCtx.TTowerEquipmentModel.FindListByTowerId(l.ctx, req.Id)
	if err != nil {
		return resp, err
	}

	if len(towerEquipmentRelateIDList) == 0 {
		return resp, nil
	}

	equipmentIDList := make([]int64, 0)
	for _, v := range towerEquipmentRelateIDList {
		equipmentIDList = append(equipmentIDList, v.EquipmentId)
	}
	// 然后查具体的设备表
	equipmentList, err := l.svcCtx.TEquipmentDetailModel.FindByIds(l.ctx, equipmentIDList)
	if err != nil {
		return resp, err
	}

	for _, v := range equipmentList {
		respEquipmentList = append(respEquipmentList, types.EquipmentDetail{
			Id:            v.Id,
			Name:          v.Name,
			Type:          v.Type,
			SpecificModel: v.SpecificModel.String,
			Unit:          v.Unit.String,
			Status:        v.Status.String,
			Image:         v.Image.String,
			CreateUserId:  v.CreateUserId.Int64,
		})
	}

	// 组合返回
	resp.PageData = respEquipmentList
	resp.Total = int64(len(respEquipmentList))

	return
}
