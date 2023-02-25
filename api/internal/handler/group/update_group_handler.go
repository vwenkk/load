package group

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/vwenkk/load/api/internal/logic/group"
	"github.com/vwenkk/load/api/internal/svc"
	"github.com/vwenkk/load/api/internal/types"
)

// swagger:route post /group/update group UpdateGroup
//
// Update Group information | 更新分组
//
// Update Group information | 更新分组
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: GroupInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewUpdateGroupLogic(r, svcCtx)
		resp, err := l.UpdateGroup(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
