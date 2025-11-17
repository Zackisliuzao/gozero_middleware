package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

// 全局中间件
func MiddlewareDemoFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("request ... ")
		next(w, r)
		logx.Info("reponse ... ")
	}
}
