package equipment

import (
	"net/http"

	"be/api/internal/logic/equipment"
	"be/api/internal/svc"
	"be/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateEquipmentClassHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateEquipmentClassReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := equipment.NewUpdateEquipmentClassLogic(r.Context(), svcCtx)
		resp, err := l.UpdateEquipmentClass(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
