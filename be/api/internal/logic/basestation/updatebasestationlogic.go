package basestation

import (
	"context"
	"database/sql"

	"be/api/internal/svc"
	"be/api/internal/types"
	"be/model/mysql"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBaseStationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBaseStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBaseStationLogic {
	return &UpdateBaseStationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBaseStationLogic) UpdateBaseStation(req *types.UpdateBaseStationReq) (resp *types.UpdateBaseStationResp, err error) {
	data := mysql.TBaseStation{
		Name:    req.Name,
		Address: req.Address,
		Image:   sql.NullString{String: req.Image, Valid: true},
		// PlanTime: sql.NullTime{String: req.PlanTime, Valid: true},
	}

	err = l.svcCtx.TBaseStation.Update(l.ctx, &data)
	if err != nil {
		return resp, err
	}

	return resp, err
}
