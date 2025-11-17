// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"greet/internal/config"

	"greet/internal/middleware"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config           config.Config
	GreetMiddleware1 rest.Middleware
	GreetMiddleware2 rest.Middleware
	RedisClient      *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		GreetMiddleware1: middleware.NewGreetMiddleware1Middleware().Handle,
		GreetMiddleware2: middleware.NewGreetMiddleware2Middleware().Handle,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
		}),
	}
}
