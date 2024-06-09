package tower

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTowerDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTowerDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTowerDetailsLogic {
	return &ListTowerDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTowerDetailsLogic) ListTowerDetails(req *types.ListTowerDetailsReq) (resp *types.ListTowerDetailsResp, err error) {

	dataList, err := l.svcCtx.TTowerDetailModel.List(l.ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	var towerDetails []types.TowerDetail
	for _, data := range dataList {
		towerDetail := types.TowerDetail{
			Id:          data.Id,
			Name:        data.Name,
			SubitemId:   data.SubitemId,
			Address:     data.Address,
			Image:       data.Image.String,
			CheckStatus: data.CheckStatus.String,
			CheckTime:   data.CheckTime.Time.String(),
			CheckUserId: data.CheckUserId.Int64,
			PrincipalId: data.PrincipalId.Int64,
			PlanTime:    data.PlanTime.Time.Format("2006-01-02 15:04:05"),
			CreateTime:  data.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:  data.UpdateTime.Format("2006-01-02 15:04:05"),
		}
		towerDetails = append(towerDetails, towerDetail)
	}
	resp = &types.ListTowerDetailsResp{
		TowerDetails: towerDetails,
	}

	return
}
