package user

import (
	"context"
	"encoding/json"
	"fmt"

	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListReq) (resp *types.UserListRes, err error) {
	jsonStr := `{
		"pageData": [
			{
				"id": 1,
				"username": "admin",
				"enable": true,
				"createTime": "2023-11-18T08:18:59.150Z",
				"updateTime": "2023-11-18T08:18:59.150Z",
				"roles": [
					{
						"id": 1,
						"code": "SUPER_ADMIN",
						"name": "超级管理员",
						"enable": true
					},
					{
						"id": 2,
						"code": "ROLE_QA",
						"name": "质检员",
						"enable": true
					}
				],
				"gender": null,
				"avatar": "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80",
				"address": null,
				"email": null
			}
		],
		"total": 1
	}`

	err = json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	return
}
