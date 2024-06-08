package equipmentdetail

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEquipmentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEquipmentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEquipmentDetailLogic {
	return &GetEquipmentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEquipmentDetailLogic) GetEquipmentDetail(req *types.GetEquipmentDetailReq) (resp *types.GetEquipmentDetailResp, err error) {

	equipmentDetail, err := l.svcCtx.TEquipmentDetailModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	respEquipmentDetail := &types.EquipmentDetail{
		Id:            equipmentDetail.Id,
		ClassId:       equipmentDetail.ClassId,
		Name:          equipmentDetail.Name,
		Type:          equipmentDetail.Type,
		SpecificModel: equipmentDetail.SpecificModel.String,
		Unit:          equipmentDetail.Unit.String,
		Status:        equipmentDetail.Status.String,
		Image:         equipmentDetail.Image.String,
		CreateUserId:  equipmentDetail.CreateUserId.Int64,
	}

	resp = &types.GetEquipmentDetailResp{
		EquipmentDetail: *respEquipmentDetail,
	}

	return resp, nil
}
