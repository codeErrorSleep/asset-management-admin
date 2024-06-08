package equipmentclass

import (
	"be/api/internal/logic/equipmentclass"
	"be/api/internal/svc"
	"be/api/internal/types"
	"be/api/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteEquipmentClassHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteEquipmentClassReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := equipmentclass.NewDeleteEquipmentClassLogic(r.Context(), svcCtx)
		resp, err := l.DeleteEquipmentClass(&req)
		response.Response(w, resp, err)
	}
}
