package towerequipment

import (
	"be/api/internal/logic/towerequipment"
	"be/api/internal/svc"
	"be/api/internal/types"
	"be/api/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateTowerEquipmentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateTowerEquipmentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := towerequipment.NewCreateTowerEquipmentLogic(r.Context(), svcCtx)
		resp, err := l.CreateTowerEquipment(&req)
		response.Response(w, resp, err)
	}
}
