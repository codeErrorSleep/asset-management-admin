package equipment

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEquipmentClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateEquipmentClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEquipmentClassLogic {
	return &CreateEquipmentClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateEquipmentClassLogic) CreateEquipmentClass(req *types.CreateEquipmentClassReq) (resp *types.CreateEquipmentClassResp, err error) {
	// todo: add your logic here and delete this line

	return
}
