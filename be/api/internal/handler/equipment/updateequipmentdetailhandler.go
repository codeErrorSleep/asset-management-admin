package equipment

import (
	"net/http"

	"be/api/internal/logic/equipment"
	"be/api/internal/svc"
	"be/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateEquipmentDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateEquipmentDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := equipment.NewUpdateEquipmentDetailLogic(r.Context(), svcCtx)
		resp, err := l.UpdateEquipmentDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
