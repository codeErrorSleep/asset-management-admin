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

	data, err := l.svcCtx.TBaseStation.ListByPage(l.ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	var baseStations []types.BaseStation
	for _, item := range data {
		baseStations = append(baseStations, types.BaseStation{
			Id:          item.Id,
			Name:        item.Name,
			Address:     item.Address,
			Image:       item.Image.String,
			CheckStatus: item.CheckStatus.String,
			CheckTime:   item.CheckTime.Time.Format("2006-01-02 15:04:05"),
			CheckUserId: item.CheckUserId.Int64,
			PrincipalId: item.PrincipalId.Int64,
			PlanTime:    item.PlanTime.Time.Format("2006-01-02 15:04:05"),
			CreateTime:  item.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:  item.UpdateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return &types.ListBaseStationsResp{
		BaseStations: baseStations,
		Total:        int64(len(data)),
	}, nil
}
