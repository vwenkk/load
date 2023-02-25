package group

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/vwenkk/load/api/internal/logic/group"
	"github.com/vwenkk/load/api/internal/svc"
	"github.com/vwenkk/load/api/internal/types"
)

// swagger:route post /group/list group GetGroupList
//
// Get Group list | 获取分组列表
//
// Get Group list | 获取分组列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: GroupListReq
//
// Responses:
//  200: GroupListResp

func GetGroupListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewGetGroupListLogic(r, svcCtx)
		resp, err := l.GetGroupList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
