package equipmentclass

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEquipmentClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEquipmentClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEquipmentClassLogic {
	return &GetEquipmentClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEquipmentClassLogic) GetEquipmentClass(req *types.GetEquipmentClassReq) (resp *types.GetEquipmentClassResp, err error) {
	// todo: add your logic here and delete this line

	return
}
