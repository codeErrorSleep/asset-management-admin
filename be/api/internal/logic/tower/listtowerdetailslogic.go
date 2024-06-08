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
	// todo: add your logic here and delete this line

	return
}
