package group

import (
	"context"
	"github.com/vwenkk/load/rpc/loadclient"
	"net/http"

	"github.com/vwenkk/load/api/internal/svc"
	"github.com/vwenkk/load/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewUpdateGroupLogic(r *http.Request, svcCtx *svc.ServiceContext) *UpdateGroupLogic {
	return &UpdateGroupLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *UpdateGroupLogic) UpdateGroup(req *types.GroupInfo) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.LoadRpc.UpdateGroup(l.ctx, &loadclient.GroupInfo{
		Id:     req.Id,
		Name:   req.Name,
		Remark: req.Remark,
	})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, result.Msg)}, nil

}
