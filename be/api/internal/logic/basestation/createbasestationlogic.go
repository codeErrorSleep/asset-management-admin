package basestation

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

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

	return
}
