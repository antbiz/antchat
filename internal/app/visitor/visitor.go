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
		group.ALL("/widget", api.Widget.Index)
		group.POST("/signin", api.Visitor.Signin)

		group.Middleware(middleware.Auth, middleware.CtxVisitor)

		group.ALL("/chat", ws.VisitorChatHandler)
		group.POST("/send", api.Msg.Send)
		group.GET("/history", api.Msg.History)
	})
}
