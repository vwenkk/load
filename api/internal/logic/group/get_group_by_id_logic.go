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

type GetGroupByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetGroupByIdLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetGroupByIdLogic {
	return &GetGroupByIdLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetGroupByIdLogic) GetGroupById(req *types.IDReq) (resp *types.GroupInfoResp, err error) {
	data, err := l.svcCtx.LoadRpc.GetGroupById(l.ctx, &loadclient.IDReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.GroupInfoResp{BaseDataInfo: types.BaseDataInfo{
		Code: 0,
		Msg:  l.svcCtx.Trans.Trans(l.lang, i18n.Success),
	}}
	resp.Data = types.GroupInfo{
		BaseInfo: types.BaseInfo{
			Id:        data.GetId(),
			CreatedAt: data.GetCreatedAt(),
			UpdatedAt: data.GetUpdatedAt(),
		},
		Name:   data.GetName(),
		Remark: data.GetRemark(),
	}
	return resp, nil
}
