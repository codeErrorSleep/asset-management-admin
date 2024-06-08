package equipmentdetail

import (
	"context"
	"database/sql"

	"be/api/internal/svc"
	"be/api/internal/types"
	"be/model/mysql"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEquipmentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateEquipmentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEquipmentDetailLogic {
	return &UpdateEquipmentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateEquipmentDetailLogic) UpdateEquipmentDetail(req *types.UpdateEquipmentDetailReq) (resp *types.UpdateEquipmentDetailResp, err error) {

	updateDetail := mysql.TEquipmentDetail{
		Id:            req.Id,
		ClassId:       req.ClassId,
		Name:          req.Name,
		Type:          req.Type,
		SpecificModel: sql.NullString{String: req.SpecificModel, Valid: true},
		Unit:          sql.NullString{String: req.Unit, Valid: true},
		Status:        sql.NullString{String: req.Status, Valid: true},
		Image:         sql.NullString{String: req.Image, Valid: true},
		// CreateUserId:  sql.NullInt64{Int64: req.CreateUserId, Valid: true},
	}

	err = l.svcCtx.TEquipmentDetailModel.Update(l.ctx, &updateDetail)
	if err != nil {
		return nil, err
	}

	return
}
