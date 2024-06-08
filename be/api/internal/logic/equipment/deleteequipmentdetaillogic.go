package equipment

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEquipmentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteEquipmentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEquipmentDetailLogic {
	return &DeleteEquipmentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteEquipmentDetailLogic) DeleteEquipmentDetail(req *types.DeleteEquipmentDetailReq) (resp *types.DeleteEquipmentDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
