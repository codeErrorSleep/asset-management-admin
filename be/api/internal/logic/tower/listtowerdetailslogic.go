package tower

import (
	"context"
	"encoding/json"
	"strconv"

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
	pageNo, err := strconv.ParseInt(req.PageNo, 10, 64)
	if err != nil {
		return nil, err
	}
	pageSize, err := strconv.ParseInt(req.PageSize, 10, 64)
	if err != nil {
		return nil, err
	}

	dataList, err := l.svcCtx.TTowerDetailModel.List(l.ctx, pageNo, pageSize)
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
			CheckStatus: data.CheckStatus.String,
			CheckTime:   data.CheckTime.Time.String(),
			CheckUserId: data.CheckUserId.Int64,
			PrincipalId: data.PrincipalId.Int64,
			PlanTime:    data.PlanTime.Time.Format("2006-01-02 15:04:05"),
			CreateTime:  data.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:  data.UpdateTime.Format("2006-01-02 15:04:05"),
		}
		towerDetail.Image = make([]string, 0)
		if len(data.Image.String) > 0 {
			err = json.Unmarshal([]byte(data.Image.String), &towerDetail.Image)
			logx.Error(err)
		}

		towerDetails = append(towerDetails, towerDetail)
	}
	resp = &types.ListTowerDetailsResp{
		PageData: towerDetails,
	}

	return
}
