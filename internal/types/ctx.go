package types

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Context struct {
	Session *ghttp.Session
	User    *ContextUser
	Data    g.Map // 自定KV变量，业务模块根据需要设置，不固定
}

// ContextUser 请求上下文中的用户信息
type ContextUser struct {
	ID       string `json:"_id"`      // 用户id
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
	Phone    string `json:"phone"`    // 手机号
	Email    string `json:"email"`    // 邮箱
	Avatar   string `json:"avatar"`   // 头像
	Language string `json:"language"` // 语言
	Role     int    `json:"role"`     // 角色
	Sid      string `json:"sid"`      // session id
}
