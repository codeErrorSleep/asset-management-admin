package basestation

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBaseStationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBaseStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBaseStationLogic {
	return &DeleteBaseStationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBaseStationLogic) DeleteBaseStation(req *types.DeleteBaseStationReq) (resp *types.DeleteBaseStationResp, err error) {
	// todo: add your logic here and delete this line

	err = l.svcCtx.TBaseStation.Delete(l.ctx, req.Id)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
