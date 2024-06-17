package tower

import (
	"context"
	"encoding/json"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTowerDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTowerDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTowerDetailLogic {
	return &GetTowerDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTowerDetailLogic) GetTowerDetail(req *types.GetTowerDetailReq) (resp *types.GetTowerDetailResp, err error) {

	towerDetail, err := l.svcCtx.TTowerDetailModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	respTowerDetail := types.TowerDetail{
		Id:          towerDetail.Id,
		SubitemId:   towerDetail.SubitemId,
		Name:        towerDetail.Name,
		Address:     towerDetail.Address,
		CheckStatus: towerDetail.CheckStatus.String,
		CheckTime:   towerDetail.CheckTime.Time.Format("2006-01-02 15:04:05"),
		CheckUserId: towerDetail.CheckUserId.Int64,
		PrincipalId: towerDetail.PrincipalId.Int64,
		PlanTime:    towerDetail.PlanTime.Time.Format("2006-01-02 15:04:05"),
		CreateTime:  towerDetail.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:  towerDetail.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	respTowerDetail.Image = make([]string, 0)
	if len(towerDetail.Image.String) > 0 {
		err = json.Unmarshal([]byte(towerDetail.Image.String), &respTowerDetail.Image)
		logx.Error(err)
	}

	resp = &types.GetTowerDetailResp{
		TowerDetail: respTowerDetail,
	}

	return resp, nil
}
