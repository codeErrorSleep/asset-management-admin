package user

import (
	"be/api/internal/logic/user"
	"be/api/internal/svc"
	"be/api/response"
	"net/http"
)

func UserLogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserLogoutLogic(r.Context(), svcCtx)
		err := l.UserLogout()
		response.Response(w, nil, err)
	}
}
