package equipmentclass

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEquipmentClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteEquipmentClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEquipmentClassLogic {
	return &DeleteEquipmentClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteEquipmentClassLogic) DeleteEquipmentClass(req *types.DeleteEquipmentClassReq) (resp *types.DeleteEquipmentClassResp, err error) {
	// todo: add your logic here and delete this line

	return
}
