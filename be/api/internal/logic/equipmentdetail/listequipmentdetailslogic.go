package equipmentdetail

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListEquipmentDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListEquipmentDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListEquipmentDetailsLogic {
	return &ListEquipmentDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListEquipmentDetailsLogic) ListEquipmentDetails(req *types.ListEquipmentDetailsReq) (resp *types.ListEquipmentDetailsResp, err error) {

	equipmentDetails, err := l.svcCtx.TEquipmentDetailModel.ListByPage(l.ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	var respEquipmentDetails []types.EquipmentDetail
	for _, equipmentDetail := range equipmentDetails {
		respEquipmentDetail := types.EquipmentDetail{
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
		respEquipmentDetails = append(respEquipmentDetails, respEquipmentDetail)
	}
	resp.EquipmentDetails = respEquipmentDetails

	return
}
