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

type UpdateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupLogic {
	return &UpdateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGroupLogic) UpdateGroup(in *load.GroupInfo) (*load.BaseResp, error) {
	err := l.svcCtx.DB.Group.UpdateOneID(in.Id).
		SetNotEmptyStatus(uint8(in.Status)).
		SetNotEmptyName(in.Name).
		SetNotEmptyRemark(in.Remark).
		Exec(l.ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
		case ent.IsConstraintError(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.UpdateFailed)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
	}

	return &load.BaseResp{Msg: i18n.UpdateSuccess}, nil

}
