package system

import (
	"github.com/antbiz/antchat/internal/app/system/api"
	"github.com/antbiz/antchat/internal/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Init() {
	s := g.Server()

	s.Group("/api/sys", func(group *ghttp.RouterGroup) {
		group.GET("/captcha", api.Captcha.Generate)
		group.POST("/account/signin", api.User.SigninByAccount)

		group.Middleware(middleware.Auth)
		group.GET("/account/info", api.User.GetInfo)

		group.GET("/visitor/{id}", api.Visitor.Get)
		group.PUT("/visitor/{id}", api.Visitor.Update)
	})
}
