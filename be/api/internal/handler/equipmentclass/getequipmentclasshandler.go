package equipmentclass

import (
	"net/http"

	"be/api/internal/logic/equipmentclass"
	"be/api/internal/svc"
	"be/api/internal/types"
	"be/api/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetEquipmentClassHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetEquipmentClassReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := equipmentclass.NewGetEquipmentClassLogic(r.Context(), svcCtx)
		resp, err := l.GetEquipmentClass(&req)
		response.Response(w, resp, err)
	}
}
