package api

import (
	"github.com/antbiz/antchat/internal/app/system/dto"
	"github.com/antbiz/antchat/internal/pkg/captcha"
	"github.com/antbiz/antchat/internal/pkg/resp"
	"github.com/gogf/gf/net/ghttp"
)

// Captcha 验证码接口
var Captcha = &captchaApi{}

type captchaApi struct{}

// Generate .
func (captchaApi) Generate(r *ghttp.Request) {
	id, b64s, err := captcha.Generate()
	if err != nil {
		resp.InternalServer(r, "err_gen_captcha", "生成验证码失败", err)
	}
	resp.OK(r, &dto.CaptchaGenReply{
		ID:     id,
		Base64: b64s,
	})
}
