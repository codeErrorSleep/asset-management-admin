package user

import (
	"be/api/internal/logic/user"
	"be/api/internal/svc"
	"be/api/internal/types"
	"be/api/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func RoleRemoveUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PatchRoleOperateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewRoleRemoveUserLogic(r.Context(), svcCtx)
		err := l.RoleRemoveUser(&req)
		response.Response(w, nil, err)
	}
}
