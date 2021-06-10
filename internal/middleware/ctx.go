package middleware

import (
	"github.com/antbiz/antchat/internal/shared"
	"github.com/antbiz/antchat/internal/types"
	"github.com/gogf/gf/net/ghttp"
)

// CtxUser 系统用户自定义上下文变量
func CtxUser(r *ghttp.Request) {
	if sessionUserID := r.Session.GetString("id"); sessionUserID != "" {
		shared.Ctx.InitCtxUser(r, &types.ContextUser{
			ID:       sessionUserID,
			Username: r.Session.GetString("username"),
			Phone:    r.Session.GetString("phone"),
			Email:    r.Session.GetString("email"),
			Avatar:   r.Session.GetString("avatar"),
			Language: r.Session.GetString("language"),
			Role:     r.Session.GetInt("role"),
			Sid:      r.Session.GetString("sid"),
		})
	}

	r.Middleware.Next()
}

// CtxUser 访客自定义上下文变量
func CtxVisitor(r *ghttp.Request) {
	if sessionUserID := r.Session.GetString("id"); sessionUserID != "" {
		shared.Ctx.InitCtxVisitor(r, &types.ContextVisitor{
			ID:       sessionUserID,
			AgentID:  r.Session.GetString("agentID"),
			Nickname: r.Session.GetString("nickname"),
			Language: r.Session.GetString("language"),
		})
	}

	r.Middleware.Next()
}
