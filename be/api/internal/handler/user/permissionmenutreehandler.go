package user

import (
	"be/api/internal/logic/user"
	"be/api/internal/svc"
	"be/api/response"
	"net/http"
)

func PermissionMenuTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewPermissionMenuTreeLogic(r.Context(), svcCtx)
		err := l.PermissionMenuTree()
		response.Response(w, nil, err)
	}
}
