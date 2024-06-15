package user

import (
	"be/api/internal/logic/user"
	"be/api/internal/svc"
	"be/api/response"
	"net/http"
)

func PermissionDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewPermissionDeleteLogic(r.Context(), svcCtx)
		err := l.PermissionDelete()
		response.Response(w, nil, err)
	}
}
