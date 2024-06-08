package tower

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTowerDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTowerDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTowerDetailLogic {
	return &DeleteTowerDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTowerDetailLogic) DeleteTowerDetail(req *types.DeleteTowerDetailReq) (resp *types.DeleteTowerDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
