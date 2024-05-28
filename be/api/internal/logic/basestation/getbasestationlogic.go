package basestation

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBaseStationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBaseStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBaseStationLogic {
	return &GetBaseStationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBaseStationLogic) GetBaseStation(req *types.GetBaseStationReq) (resp *types.GetBaseStationResp, err error) {
	baseStation, err := l.svcCtx.TBaseStation.FindOne(l.ctx, req.Id)
	if err != nil {
		return resp, nil
	}

	resp = &types.GetBaseStationResp{
		BaseStation: types.BaseStation{
			Id:         baseStation.Id,
			Name:       baseStation.Name,
			CreateTime: baseStation.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: baseStation.UpdateTime.Format("2006-01-02 15:04:05"),
		},
	}

	return resp, nil
}
