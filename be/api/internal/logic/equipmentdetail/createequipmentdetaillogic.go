package equipmentdetail

import (
	"context"
	"database/sql"
	"errors"

	"be/api/internal/svc"
	"be/api/internal/types"
	"be/model/mysql"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEquipmentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateEquipmentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEquipmentDetailLogic {
	return &CreateEquipmentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateEquipmentDetailLogic) CreateEquipmentDetail(req *types.CreateEquipmentDetailReq) (resp *types.CreateEquipmentDetailResp, err error) {
	data := mysql.TEquipmentDetail{
		ClassId:       req.ClassId,
		Name:          req.Name,
		Type:          req.Type,
		SpecificModel: sql.NullString{String: req.SpecificModel, Valid: true},
		Unit:          sql.NullString{String: req.Unit, Valid: true},
		Status:        sql.NullString{String: req.Status, Valid: true},
		Image:         sql.NullString{String: req.Image, Valid: true},
		CreateUserId:  sql.NullInt64{Int64: req.CreateUserId, Valid: true},
	}

	sqlResult, err := l.svcCtx.TEquipmentDetailModel.Insert(l.ctx, &data)
	if err != nil {
		return nil, err
	}

	if rowAffected, _ := sqlResult.RowsAffected(); rowAffected == 0 {
		return nil, errors.New("insert failed")
	}

	return
}
