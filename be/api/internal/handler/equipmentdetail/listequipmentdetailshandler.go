package equipmentdetail

import (
	"be/api/internal/logic/equipmentdetail"
	"be/api/internal/svc"
	"be/api/internal/types"
	"be/api/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func ListEquipmentDetailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListEquipmentDetailsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := equipmentdetail.NewListEquipmentDetailsLogic(r.Context(), svcCtx)
		resp, err := l.ListEquipmentDetails(&req)
		response.Response(w, resp, err)
	}
}
