package system

import (
	"github.com/antbiz/antchat/internal/app/system/api"
	"github.com/antbiz/antchat/internal/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Init() {
	s := g.Server()

	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.POST("/account/signin", api.User.SigninByAccount)

		group.Middleware(middleware.Auth)
		group.GET("/account/info", api.User.GetInfo)
	})
}
