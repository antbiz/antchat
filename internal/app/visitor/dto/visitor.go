package dto

type VisitorSigninReq struct {
	Captcha   string `json:"captcha" v:"required"`
	Domain    string `json:"domain" v:"required"`
	Geo       string `json:"geo"`
	VisitorID string `json:"visitorID"`
}
