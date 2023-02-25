package group

import (
	"context"
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
	// todo: add your logic here and delete this line

	return
}
