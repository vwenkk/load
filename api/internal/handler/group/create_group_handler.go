package group

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/vwenkk/load/api/internal/logic/group"
	"github.com/vwenkk/load/api/internal/svc"
	"github.com/vwenkk/load/api/internal/types"
)

// swagger:route post /group/create group CreateGroup
//
// Create Group information | 创建分组
//
// Create Group information | 创建分组
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: GroupInfo
//
// Responses:
//  200: BaseMsgResp

func CreateGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewCreateGroupLogic(r, svcCtx)
		resp, err := l.CreateGroup(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
