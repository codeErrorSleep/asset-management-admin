package tower

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
