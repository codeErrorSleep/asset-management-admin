package tower

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTowerDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTowerDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTowerDetailLogic {
	return &CreateTowerDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTowerDetailLogic) CreateTowerDetail(req *types.CreateTowerDetailReq) (resp *types.CreateTowerDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
