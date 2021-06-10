package agent

import (
	"github.com/antbiz/antchat/internal/app/agent/api"
	"github.com/antbiz/antchat/internal/app/ws"
	"github.com/antbiz/antchat/internal/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Init() {
	s := g.Server()

	s.Group("/api/v1/agent", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth, middleware.CtxUser)

		group.GET("/chat", ws.AgentChatHandler)
		group.POST("/send", api.Msg.Send)
		group.GET("/history", api.Msg.History)
		group.GET("/conversations", api.Msg.Conversations)
	})
}
