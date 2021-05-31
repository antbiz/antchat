package api

import "github.com/gogf/gf/net/ghttp"

var Visitor = new(visitorApi)

type visitorApi struct{}

// Register 访客注册
func (visitorApi) Register(r *ghttp.Request) {

}
