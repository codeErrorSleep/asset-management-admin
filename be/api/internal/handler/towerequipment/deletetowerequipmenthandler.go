package towerequipment

import (
	"net/http"

	"be/api/internal/logic/towerequipment"
	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
)

func DeleteTowerEquipmentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteTowerEquipmentReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := towerequipment.NewDeleteTowerEquipmentLogic(r.Context(), svcCtx)
		resp, err := l.DeleteTowerEquipment(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
