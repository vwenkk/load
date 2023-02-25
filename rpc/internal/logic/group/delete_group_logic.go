package group

import (
	"context"

	"github.com/vwenkk/load/rpc/internal/svc"
	"github.com/vwenkk/load/rpc/types/load"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupLogic {
	return &DeleteGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteGroupLogic) DeleteGroup(in *load.IDsReq) (*load.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &load.BaseResp{}, nil
}
