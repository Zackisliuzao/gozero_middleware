// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"flag"
	"fmt"

	"greet/internal/config"
	"greet/internal/handler"
	"greet/internal/middleware"
	"greet/internal/svc"

	paramsvalidator "github.com/lerity-yao/param-validator"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/greet-api.yaml", "the config file")

func main() {
	// 1. 创建自定义校验器
	vd := paramsvalidator.MustNewHttpxParseValidator(paramsvalidator.Conf{
		ZhTrans: true, // 是否使用中文提示，默认就是 true
	})

	// 2. 替换 httpx 的全局校验器
	httpx.SetValidator(vd)

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	// 使用全局中间件
	server.Use(middleware.MiddlewareDemoFunc)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
