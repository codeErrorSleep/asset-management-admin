package tower

import (
	"context"
	"errors"

	"be/api/internal/svc"
	"be/api/internal/types"
	"be/model/mysql"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTowerDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTowerDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTowerDetailLogic {
	return &CreateTowerDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTowerDetailLogic) CreateTowerDetail(req *types.CreateTowerDetailReq) (resp *types.CreateTowerDetailResp, err error) {
	insertData := mysql.TTowerDetail{
		SubitemId: req.SubitemId,
		Name:      req.Name,
		Address:   req.Address,
	}

	sqlResult, err := l.svcCtx.TTowerDetailModel.Insert(l.ctx, &insertData)
	if err != nil {
		return nil, err
	}

	if attect, _ := sqlResult.RowsAffected(); attect == 0 {
		return nil, errors.New("insert failed")
	}

	return resp, nil
}
