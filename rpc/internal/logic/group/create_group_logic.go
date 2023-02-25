package group

import (
	"context"

	"github.com/vwenkk/load/rpc/internal/svc"
	"github.com/vwenkk/load/rpc/types/load"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Group management
func (l *CreateGroupLogic) CreateGroup(in *load.GroupInfo) (*load.BaseIDResp, error) {
	// todo: add your logic here and delete this line

	return &load.BaseIDResp{}, nil
}
