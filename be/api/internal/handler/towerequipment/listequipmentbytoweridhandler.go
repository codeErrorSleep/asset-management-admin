package towerequipment

import (
	"net/http"

	"be/api/internal/logic/towerequipment"
	"be/api/internal/svc"
	"be/api/internal/types"
	"be/api/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListEquipmentByTowerIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListEquipmentByTowerIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := towerequipment.NewListEquipmentByTowerIdLogic(r.Context(), svcCtx)
		resp, err := l.ListEquipmentByTowerId(&req)
		response.Response(w, resp, err)
	}
}
