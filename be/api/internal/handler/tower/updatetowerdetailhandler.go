package tower

import (
	"net/http"

	"be/api/internal/logic/tower"
	"be/api/internal/svc"
	"be/api/internal/types"
	"be/api/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateTowerDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateTowerDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tower.NewUpdateTowerDetailLogic(r.Context(), svcCtx)
		resp, err := l.UpdateTowerDetail(&req)
		response.Response(w, resp, err)
	}
}
