package types

// ContextUser 请求上下文中的用户信息
type ContextUser struct {
	ID       string `json:"id"`       // 用户id
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
	Phone    string `json:"phone"`    // 手机号
	Email    string `json:"email"`    // 邮箱
	Avatar   string `json:"avatar"`   // 头像
	Language string `json:"language"` // 语言
	Role     int    `json:"role"`     // 角色
	Sid      string `json:"sid"`      // session id
}

// ContextVisitor 请求上下文中的访客信息
type ContextVisitor struct {
	ID       string `json:"id"`       // 访客id
	AgentID  string `json:"agentID"`  // 对接此访客的客服id
	Nickname string `json:"nickname"` // 昵称
	Language string `json:"language"` // 语言
}
