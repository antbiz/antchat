package dto

// VisitorUpdateReq 修改访客信息请求
type VisitorUpdateReq struct {
	Nickname string `v:"required"`
	Phone    string `v:"phone"`
	Email    string `v:"email"`
}
