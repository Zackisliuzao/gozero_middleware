// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetMiddleware1Middleware struct {
}

func NewGreetMiddleware1Middleware() *GreetMiddleware1Middleware {
	return &GreetMiddleware1Middleware{}
}

func (m *GreetMiddleware1Middleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		logx.Info("greetMiddleware1 request ... ")
		// Passthrough to next handler if need
		next(w, r)
		logx.Info("greetMiddleware1 reponse ... ")
	}
}
