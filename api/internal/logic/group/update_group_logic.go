package group

import (
	"context"
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
	// todo: add your logic here and delete this line

	return
}
