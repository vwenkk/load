package group

import (
	"context"
	"github.com/vwenkk/load/pkg/statuserr"
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
	result, err := l.svcCtx.DB.Group.Get(l.ctx, in.GetId())
	if err != nil {
		return nil, statuserr.Wrap(err, in)
	}
	return &load.GroupInfo{
		Id:        result.ID,
		CreatedAt: result.CreatedAt.Unix(),
		UpdatedAt: result.UpdatedAt.Unix(),
		Status:    uint32(result.Status),
		Name:      result.Name,
		Remark:    result.Remark,
	}, nil
}
