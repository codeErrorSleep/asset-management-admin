package user

import (
	"be/api/internal/logic/user"
	"be/api/internal/svc"
	"be/api/response"
	"net/http"
)

func UserUpdatePasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserUpdatePasswordLogic(r.Context(), svcCtx)
		err := l.UserUpdatePassword()
		response.Response(w, nil, err)
	}
}
