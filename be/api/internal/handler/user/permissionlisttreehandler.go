package user

import (
	"be/api/internal/logic/user"
	"be/api/internal/svc"
	"be/api/response"
	"net/http"
)

func PermissionListTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewPermissionListTreeLogic(r.Context(), svcCtx)
		err := l.PermissionListTree()
		response.Response(w, nil, err)
	}
}
