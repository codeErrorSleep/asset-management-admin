package user

import (
	"context"

	"be/api/internal/svc"
	"be/api/internal/types"
	"be/model/mysql"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddLogic) UserAdd(req *types.AddUserReq) error {

	insertData := mysql.User{
		Username: req.Username,
		Password: req.Password,
	}

	_, err := l.svcCtx.TUser.Insert(l.ctx, &insertData)
	if err != nil {
		return err
	}

	return nil
}
