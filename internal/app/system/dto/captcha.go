package dto

// CaptchaGenReq 获取验证码响应
type CaptchaGenReply struct {
	ID     string `json:"id"`
	Base64 string `json:"base64"`
}
