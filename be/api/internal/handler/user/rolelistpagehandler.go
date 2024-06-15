package user

import (
	"be/api/internal/logic/user"
	"be/api/internal/svc"
	"be/api/response"
	"net/http"
)

func RoleListPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewRoleListPageLogic(r.Context(), svcCtx)
		err := l.RoleListPage()
		response.Response(w, nil, err)
	}
}
