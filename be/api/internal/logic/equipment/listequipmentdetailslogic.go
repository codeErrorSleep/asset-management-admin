package equipment

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListEquipmentDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListEquipmentDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListEquipmentDetailsLogic {
	return &ListEquipmentDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListEquipmentDetailsLogic) ListEquipmentDetails(req *types.ListEquipmentDetailsReq) (resp *types.ListEquipmentDetailsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
