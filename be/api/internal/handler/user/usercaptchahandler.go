package user

import (
	"be/api/internal/logic/user"
	"be/api/internal/svc"
	"be/api/response"
	"net/http"
)

func UserCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserCaptchaLogic(r.Context(), svcCtx)
		err := l.UserCaptcha()
		response.Response(w, nil, err)
	}
}
