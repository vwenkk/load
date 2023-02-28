package group

import (
	"context"
	"github.com/vwenkk/load/pkg/ent"
	"github.com/vwenkk/load/pkg/ent/group"
	"github.com/vwenkk/load/pkg/i18n"
	"github.com/vwenkk/load/pkg/msg/logmsg"
	"github.com/vwenkk/load/pkg/statuserr"

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

	_, err := l.svcCtx.DB.Group.Delete().Where(group.IDIn(in.Ids...)).Exec(l.ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
	}

	return &load.BaseResp{Msg: i18n.DeleteSuccess}, nil

}
