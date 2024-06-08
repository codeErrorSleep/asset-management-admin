package equipmentclass

import (
	"context"
	"database/sql"
	"errors"

	"be/api/internal/svc"
	"be/api/internal/types"
	"be/model/mysql"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEquipmentClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateEquipmentClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEquipmentClassLogic {
	return &CreateEquipmentClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateEquipmentClassLogic) CreateEquipmentClass(req *types.CreateEquipmentClassReq) (resp *types.CreateEquipmentClassResp, err error) {
	insertData := mysql.TEquipmentClass{
		Name:   req.Name,
		Code:   req.Code,
		Status: int64(req.Status),
		PId:    sql.NullInt64{Int64: req.PId, Valid: true},
	}

	sqlResult, err := l.svcCtx.TEquipmentClassModel.Insert(l.ctx, &insertData)
	if err != nil {
		return nil, err
	}

	if rowAffected, _ := sqlResult.RowsAffected(); rowAffected == 0 {
		return nil, errors.New("insert failed")
	}

	return
}
