package user

import (
	"be/api/internal/logic/user"
	"be/api/internal/svc"
	"be/api/response"
	"net/http"
)

func UserDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserDeleteLogic(r.Context(), svcCtx)
		err := l.UserDelete()
		response.Response(w, nil, err)
	}
}