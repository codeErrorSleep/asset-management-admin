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
	// todo: add your logic here and delete this line

	return
}
