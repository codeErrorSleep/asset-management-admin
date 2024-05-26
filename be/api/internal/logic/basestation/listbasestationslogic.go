package basestation

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListBaseStationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListBaseStationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBaseStationsLogic {
	return &ListBaseStationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListBaseStationsLogic) ListBaseStations(req *types.ListBaseStationsReq) (resp *types.ListBaseStationsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
