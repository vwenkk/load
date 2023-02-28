package group

import (
	"context"
	"github.com/vwenkk/load/pkg/ent"
	"github.com/vwenkk/load/pkg/i18n"
	"github.com/vwenkk/load/pkg/msg/logmsg"
	"github.com/vwenkk/load/pkg/statuserr"

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

// CreateGroup Group management
func (l *CreateGroupLogic) CreateGroup(in *load.GroupInfo) (*load.BaseIDResp, error) {
	result, err := l.svcCtx.DB.Group.Create().
		SetStatus(uint8(in.Status)).
		SetName(in.Name).
		SetRemark(in.Remark).
		Save(l.ctx)
	if err != nil {
		switch {
		case ent.IsConstraintError(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.CreateFailed)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
	}

	return &load.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
