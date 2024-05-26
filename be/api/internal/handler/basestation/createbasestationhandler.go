package basestation

import (
	"net/http"

	"be/api/internal/logic/basestation"
	"be/api/internal/svc"
	"be/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateBaseStationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateBaseStationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := basestation.NewCreateBaseStationLogic(r.Context(), svcCtx)
		resp, err := l.CreateBaseStation(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
