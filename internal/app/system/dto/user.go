package dto

// UserSigninReq 用户登录请求
type UserSigninReq struct {
	Account  string `v:"required#请输入用户名/手机号/邮箱"` // 支持用户名/手机号/邮箱
	Password string `v:"required#请输入密码"`
}

// UserSignUpWithPhoneReq 用户手机号注册请求
type UserSignUpWithPhoneReq struct {
	Username  string `v:"required|length:6,12#请输入用户名|用户名长度应当在:min到:max之间"`
	Password  string `v:"required|length:6,16#请输入密码|密码长度应当在:min到:max之间"`
	Password2 string `v:"required|length:6,16|same:Password#请输入确认密码|密码长度应当在:min到:max之间|两次密码输入不相等"`
	Phone     string `v:"required|phone#请输入手机号|手机号格式不正确"`
	Captcha   string `v:"required#请输入验证码"`
}

// UserSignUpWithEmailReq 用户邮箱注册请求
type UserSignUpWithEmailReq struct {
	Username  string `v:"required|length:6,12#请输入用户名|用户名长度应当在:min到:max之间"`
	Password  string `v:"required|length:6,16#请输入密码|密码长度应当在:min到:max之间"`
	Password2 string `v:"required|length:6,16|same:Password#请输入确认密码|密码长度应当在:min到:max之间|两次密码输入不相等"`
	Email     string `v:"required|email#请输入邮箱|邮箱格式不正确"`
	Captcha   string `v:"required#请输入验证码"`
}

// UserUpdateInfoReq 用户更新个人信息请求
type UserUpdateInfoReq struct {
	Avatar   string
	Language string `v:"required|in:zh-CN,en#请选择语言|请选择中文或者英文"`
}

// UserUpdatePasswordReq 用户更新密码请求
type UserUpdatePasswordReq struct {
	OldPassword string `v:"required#请输入旧密码"`
	Password    string `v:"required|length:6,16#请输入新密码|密码长度应当在:min到:max之间"`
	Password2   string `v:"required|length:6,16|same:Password#请输入确认密码|密码长度应当在:min到:max之间|两次密码输入不相等"`
}
