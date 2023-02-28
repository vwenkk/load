package group

import (
	"context"
	"github.com/vwenkk/load/rpc/loadclient"
	"net/http"

	"github.com/vwenkk/load/api/internal/svc"
	"github.com/vwenkk/load/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewCreateGroupLogic(r *http.Request, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *CreateGroupLogic) CreateGroup(req *types.GroupInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.LoadRpc.CreateGroup(l.ctx,
		&loadclient.GroupInfo{
			Id:     req.Id,
			Name:   req.Name,
			Remark: req.Remark,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, data.Msg)}, nil
}
