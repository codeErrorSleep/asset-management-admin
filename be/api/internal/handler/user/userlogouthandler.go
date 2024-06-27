package user

import (
	"net/http"

	"be/api/internal/logic/user"
	"be/api/internal/svc"

	xhttp "github.com/zeromicro/x/http"
)

func UserLogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserLogoutLogic(r.Context(), svcCtx)
		err := l.UserLogout()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
