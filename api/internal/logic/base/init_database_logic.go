package base

import (
	"context"
	"errors"
	"github.com/vwenkk/load/pkg/enum"
	"github.com/vwenkk/load/pkg/i18n"
	"github.com/vwenkk/load/rpc/types/load"
	"github.com/zeromicro/go-zero/core/errorx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"time"

	"github.com/vwenkk/load/api/internal/svc"
	"github.com/vwenkk/load/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewInitDatabaseLogic(r *http.Request, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.LoadRpc.InitDatabase(l.ctx, &load.Empty{})
	if err != nil && !errors.Is(err, status.Error(codes.DeadlineExceeded, "context deadline exceeded")) {
		return nil, err
	} else if errors.Is(err, status.Error(codes.DeadlineExceeded, "context deadline exceeded")) {
		for {
			// wait 10 second for initialization
			time.Sleep(time.Second * 5)
			if initState, err := l.svcCtx.Redis.Get("database_init_state"); err == nil {
				if initState == "1" {
					return nil, errorx.NewCodeError(enum.InvalidArgument,
						l.svcCtx.Trans.Trans(l.lang, i18n.AlreadyInit))
				}
			} else {
				return nil, errorx.NewCodeError(enum.Internal,
					l.svcCtx.Trans.Trans(l.lang, i18n.RedisError))
			}

			if errMsg, err := l.svcCtx.Redis.Get("database_error_msg"); err == nil {
				if errMsg != "" {
					return nil, errorx.NewCodeError(enum.Internal, errMsg)
				}
			} else {
				return nil, errorx.NewCodeError(enum.Internal,
					l.svcCtx.Trans.Trans(l.lang, i18n.RedisError))
			}
		}
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, l.svcCtx.Trans.Trans(l.lang, result.Msg))}, nil
}
