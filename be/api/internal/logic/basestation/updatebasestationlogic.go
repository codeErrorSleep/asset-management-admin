package basestation

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
