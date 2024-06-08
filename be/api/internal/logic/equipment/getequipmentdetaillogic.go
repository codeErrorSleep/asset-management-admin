package equipment

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEquipmentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEquipmentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEquipmentDetailLogic {
	return &GetEquipmentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEquipmentDetailLogic) GetEquipmentDetail(req *types.GetEquipmentDetailReq) (resp *types.GetEquipmentDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
