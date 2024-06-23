package tower

import (
	"be/api/internal/logic/tower"
	"be/api/internal/svc"
	"be/api/internal/types"
	"be/api/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func ListTowerDetailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListTowerDetailsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tower.NewListTowerDetailsLogic(r.Context(), svcCtx)
		resp, err := l.ListTowerDetails(&req)
		response.Response(w, resp, err)
	}
}
