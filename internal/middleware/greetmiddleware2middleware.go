// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetMiddleware2Middleware struct {
}

func NewGreetMiddleware2Middleware() *GreetMiddleware2Middleware {
	return &GreetMiddleware2Middleware{}
}

// 中间件2
func (m *GreetMiddleware2Middleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		logx.Info("greetMiddleware2 request ... ")
		// Passthrough to next handler if need
		next(w, r)
		logx.Info("greetMiddleware2 reponse ... ")
	}
}
