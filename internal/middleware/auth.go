package middleware

import (
	"github.com/antbiz/antchat/internal/pkg/resp"
	"github.com/gogf/gf/net/ghttp"
)

// Auth 鉴权中间件
func Auth(r *ghttp.Request) {
	if r.Session.GetString("_id") == "" {
		resp.Unauthorized(r, "unauthorized", "未登录或非法访问")
	}

	r.Middleware.Next()
}
