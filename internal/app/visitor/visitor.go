package visitor

import (
	"github.com/antbiz/antchat/internal/app/visitor/api"
	"github.com/antbiz/antchat/internal/app/ws"
	"github.com/antbiz/antchat/internal/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Init() {
	s := g.Server()

	s.Group("/api/v1/visitor", func(group *ghttp.RouterGroup) {
		group.POST("/login", api.Visitor.Login)

		group.Middleware(middleware.Auth, middleware.CtxVisitor)

		group.GET("/history", api.Msg.History)
		group.POST("/send", api.Msg.Send)
		group.GET("/chat", ws.ChatHandler)
	})
}
