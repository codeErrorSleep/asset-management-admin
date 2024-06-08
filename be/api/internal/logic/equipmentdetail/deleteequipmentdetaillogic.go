package equipmentdetail

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

	err = l.svcCtx.TEquipmentDetailModel.Delete(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return
}
