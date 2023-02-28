package base

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/vwenkk/load/pkg/i18n"
	"github.com/vwenkk/load/pkg/msg/logmsg"
	"github.com/vwenkk/load/pkg/statuserr"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/vwenkk/load/rpc/internal/svc"
	"github.com/vwenkk/load/rpc/types/load"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(in *load.Empty) (*load.BaseResp, error) {
	l.ctx = context.Background()

	// add lock to avoid duplicate initialization
	lock := redis.NewRedisLock(l.svcCtx.Redis, "init_database_lock")
	lock.SetExpire(60)
	if ok, err := lock.Acquire(); !ok || err != nil {
		if !ok {
			logx.Error("last initialization is running")
			return nil, statuserr.NewInternalError(i18n.InitRunning)
		} else {
			logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.RedisError)
		}
	}
	defer func() {
		recover()
		lock.Release()
	}()

	// initialize table structure
	if err := l.svcCtx.DB.Schema.Create(l.ctx, schema.WithForeignKeys(false), schema.WithDropColumn(true),
		schema.WithDropIndex(true)); err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, statuserr.NewInternalError(err.Error())
	}

	l.svcCtx.Redis.Setex("database_init_state", "1", 300)
	return &load.BaseResp{Msg: i18n.Success}, nil
}
