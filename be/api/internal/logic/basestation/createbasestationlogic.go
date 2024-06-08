package basestation

import (
	"context"
	"database/sql"
	"errors"

	"be/api/internal/svc"
	"be/api/internal/types"
	"be/model/mysql"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBaseStationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateBaseStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBaseStationLogic {
	return &CreateBaseStationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateBaseStationLogic) CreateBaseStation(req *types.CreateBaseStationReq) (resp *types.CreateBaseStationResp, err error) {
	// todo: add your logic here and delete this line

	data := mysql.TBaseStation{
		Name:    req.Name,
		Address: req.Address,
		Image:   sql.NullString{String: req.Image, Valid: true},
		// PlanTime: sql.NullTime{String: req.PlanTime, Valid: true},
	}

	sqlResult, err := l.svcCtx.TBaseStation.Insert(l.ctx, &data)
	if err != nil {
		return resp, err
	}

	rowAffect, err := sqlResult.RowsAffected()
	if rowAffect == 0 {
		return resp, errors.New("插入失败")
	}

	return
}
