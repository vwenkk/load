package group

import (
	"context"

	"github.com/vwenkk/load/rpc/internal/svc"
	"github.com/vwenkk/load/rpc/types/load"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupByIdLogic {
	return &GetGroupByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupByIdLogic) GetGroupById(in *load.IDReq) (*load.GroupInfo, error) {
	// todo: add your logic here and delete this line

	return &load.GroupInfo{}, nil
}
