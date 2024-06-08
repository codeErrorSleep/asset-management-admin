package tower

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
