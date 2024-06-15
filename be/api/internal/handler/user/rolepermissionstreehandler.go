package user

import (
	"be/api/internal/logic/user"
	"be/api/internal/svc"
	"be/api/response"
	"net/http"
)

func RolePermissionsTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewRolePermissionsTreeLogic(r.Context(), svcCtx)
		err := l.RolePermissionsTree()
		response.Response(w, nil, err)
	}
}
