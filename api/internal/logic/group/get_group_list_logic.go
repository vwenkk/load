package group

import (
	"context"
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
	// todo: add your logic here and delete this line

	return
}
