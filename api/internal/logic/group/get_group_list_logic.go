package group

import (
	"context"
	"github.com/vwenkk/load/pkg/i18n"
	"github.com/vwenkk/load/rpc/loadclient"
	"net/http"

	"github.com/vwenkk/load/api/internal/svc"
	"github.com/vwenkk/load/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetGroupListLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetGroupListLogic {
	return &GetGroupListLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetGroupListLogic) GetGroupList(req *types.GroupListReq) (resp *types.GroupListResp, err error) {
	list, err := l.svcCtx.LoadRpc.GetGroupList(l.ctx, &loadclient.GroupListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.GroupListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.lang, i18n.Success)
	resp.Data.Total = list.GetTotal()
	for _, v := range list.GetData() {
		resp.Data.Data = append(resp.Data.Data, types.GroupInfo{
			BaseInfo: types.BaseInfo{
				Id:        v.Id,
				CreatedAt: v.GetCreatedAt(),
				UpdatedAt: v.GetUpdatedAt(),
			},
			Name:   v.GetName(),
			Remark: v.GetRemark(),
		})
	}
	return resp, nil
}
