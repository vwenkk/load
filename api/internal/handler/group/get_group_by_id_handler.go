package group

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/vwenkk/load/api/internal/logic/group"
	"github.com/vwenkk/load/api/internal/svc"
	"github.com/vwenkk/load/api/internal/types"
)

// swagger:route post /group group GetGroupById
//
// Get Group by ID | 通过ID获取分组
//
// Get Group by ID | 通过ID获取分组
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: GroupInfoResp

func GetGroupByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewGetGroupByIdLogic(r, svcCtx)
		resp, err := l.GetGroupById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
