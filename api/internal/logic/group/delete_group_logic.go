package group

import (
	"context"
	"github.com/vwenkk/load/rpc/loadclient"
	"net/http"

	"github.com/vwenkk/load/api/internal/svc"
	"github.com/vwenkk/load/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewDeleteGroupLogic(r *http.Request, svcCtx *svc.ServiceContext) *DeleteGroupLogic {
	return &DeleteGroupLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *DeleteGroupLogic) DeleteGroup(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.LoadRpc.DeleteGroup(l.ctx, &loadclient.IDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, result.Msg)}, nil

}
