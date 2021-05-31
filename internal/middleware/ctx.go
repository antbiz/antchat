package middleware

import (
	"github.com/antbiz/antchat/internal/shared"
	"github.com/antbiz/antchat/internal/types"
	"github.com/gogf/gf/net/ghttp"
)

// Ctx 自定义上下文变量
func Ctx(r *ghttp.Request) {
	if sessionUserID := r.Session.GetString("_id"); sessionUserID != "" {
		shared.Ctx.Init(r, &types.Context{
			Session: r.Session,
			User: &types.ContextUser{
				ID:       sessionUserID,
				Username: r.Session.GetString("username"),
				Phone:    r.Session.GetString("phone"),
				Email:    r.Session.GetString("email"),
				Avatar:   r.Session.GetString("avatar"),
				Language: r.Session.GetString("language"),
				Role:     r.Session.GetInt("role"),
				Sid:      r.Session.GetString("sid"),
			},
		})
	}

	r.Middleware.Next()
}
