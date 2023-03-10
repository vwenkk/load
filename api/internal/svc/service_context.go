package svc

import (
	"github.com/vwenkk/load/api/internal/config"
	"github.com/vwenkk/load/api/internal/middleware"
	i18n2 "github.com/vwenkk/load/pkg/i18n"
	"github.com/vwenkk/load/rpc/loadclient"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/vwenkk/load/pkg/i18n"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	Casbin    *casbin.Enforcer
	Authority rest.Middleware
	Trans     *i18n.Translator
	LoadRpc   loadclient.Load
	Redis     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := redis.MustNewRedis(c.RedisConf)

	cbn := c.CasbinConf.MustNewCasbin(c.DatabaseConf.Type, c.DatabaseConf.GetDSN())

	trans := &i18n.Translator{}
	trans.NewBundle(i18n2.LocaleFS)
	trans.NewTranslator()

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds, trans).Handle,
		Trans:     trans,
		LoadRpc:   loadclient.NewLoad(zrpc.NewClientIfEnable(c.LoadRpc)),
		Redis:     rds,
	}
}
