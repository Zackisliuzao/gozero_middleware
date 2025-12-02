package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Redis redis.RedisConf
	DB    struct {
		DataSource string
	}
	Cache   cache.CacheConf
	ZhTrans bool `json:"zhTrans,optional,default=true"` // 是否开启中文, 默认中文
}
