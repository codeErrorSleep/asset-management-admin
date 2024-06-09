package tower

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"
	"be/model/mysql"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTowerDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTowerDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTowerDetailLogic {
	return &UpdateTowerDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTowerDetailLogic) UpdateTowerDetail(req *types.UpdateTowerDetailReq) (resp *types.UpdateTowerDetailResp, err error) {

	updateData := mysql.TTowerDetail{
		SubitemId: req.SubitemId,
		Name:      req.Name,
		Address:   req.Address,
	}

	err = l.svcCtx.TTowerDetailModel.Update(l.ctx, &updateData)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
