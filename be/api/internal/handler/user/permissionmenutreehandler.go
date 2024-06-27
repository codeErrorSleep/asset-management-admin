package user

import (
	"net/http"

	"be/api/internal/logic/user"
	"be/api/internal/svc"

	xhttp "github.com/zeromicro/x/http"
)

func PermissionMenuTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewPermissionMenuTreeLogic(r.Context(), svcCtx)
		err := l.PermissionMenuTree()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
