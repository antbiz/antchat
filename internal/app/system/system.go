package system

import (
	"github.com/antbiz/antchat/internal/app/system/api"
	"github.com/antbiz/antchat/internal/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Init() {
	s := g.Server()

	s.Group("/api/v1/system", func(group *ghttp.RouterGroup) {
		group.POST("/login", api.User.LoginByAccount)

		group.Middleware(middleware.Auth)
		group.GET("/account/info", api.User.GetInfo)
	})
}
