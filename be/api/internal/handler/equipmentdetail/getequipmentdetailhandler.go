package equipmentdetail

import (
	"be/api/internal/logic/equipmentdetail"
	"be/api/internal/svc"
	"be/api/internal/types"
	"be/api/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetEquipmentDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetEquipmentDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := equipmentdetail.NewGetEquipmentDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetEquipmentDetail(&req)
		response.Response(w, resp, err)
	}
}
