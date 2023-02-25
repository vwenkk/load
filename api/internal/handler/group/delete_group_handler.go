package group

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/vwenkk/load/api/internal/logic/group"
	"github.com/vwenkk/load/api/internal/svc"
	"github.com/vwenkk/load/api/internal/types"
)

// swagger:route post /group/delete group DeleteGroup
//
// Delete Group information | 删除分组信息
//
// Delete Group information | 删除分组信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewDeleteGroupLogic(r, svcCtx)
		resp, err := l.DeleteGroup(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
