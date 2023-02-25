package group

import (
	"context"
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
	// todo: add your logic here and delete this line

	return
}
