package equipment

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEquipmentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateEquipmentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEquipmentDetailLogic {
	return &UpdateEquipmentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateEquipmentDetailLogic) UpdateEquipmentDetail(req *types.UpdateEquipmentDetailReq) (resp *types.UpdateEquipmentDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
