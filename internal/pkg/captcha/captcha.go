package captcha

import "github.com/mojocn/base64Captcha"

var store = base64Captcha.DefaultMemStore

// Generate 生成 base64 验证码
func Generate() (id string, b64s string, err error) {
	drv := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	return base64Captcha.NewCaptcha(drv, store).Generate()
}

// Verify 校验验证码
func Verify(id, answer string) bool {
	return store.Verify(id, answer, true)
}
