package equipmentdetail

import (
	"net/http"

	"be/api/internal/logic/equipmentdetail"
	"be/api/internal/svc"
	"be/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
)

func CreateEquipmentDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateEquipmentDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := equipmentdetail.NewCreateEquipmentDetailLogic(r.Context(), svcCtx)
		resp, err := l.CreateEquipmentDetail(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
