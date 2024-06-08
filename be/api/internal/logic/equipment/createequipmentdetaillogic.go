package equipment

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEquipmentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateEquipmentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEquipmentDetailLogic {
	return &CreateEquipmentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateEquipmentDetailLogic) CreateEquipmentDetail(req *types.CreateEquipmentDetailReq) (resp *types.CreateEquipmentDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
