package middleware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Sid 允许通过请求参数传递sid
func Sid(r *ghttp.Request) {
	if r.GetSessionId() == "" {
		sidname := r.Server.GetSessionIdName()
		sid := r.GetString(sidname)
		r.Header.Set(sidname, sid)
		if err := r.Session.SetId(sid); err != nil {
			g.Log().Async().Errorf("middleware.Sid.SetSessionID: %v", err)
		}
	}
	r.Middleware.Next()
}
