package equipmentclass

import (
	"net/http"

	"be/api/internal/logic/equipmentclass"
	"be/api/internal/svc"
	"be/api/internal/types"
	"be/api/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateEquipmentClassHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateEquipmentClassReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := equipmentclass.NewCreateEquipmentClassLogic(r.Context(), svcCtx)
		resp, err := l.CreateEquipmentClass(&req)
		response.Response(w, resp, err)
	}
}
