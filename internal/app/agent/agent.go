package agent

import (
	"github.com/antbiz/antchat/internal/app/agent/api"
	"github.com/antbiz/antchat/internal/app/ws"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Init() {
	s := g.Server()

	s.Group("/api/v1/agent", func(group *ghttp.RouterGroup) {
		group.GET("/history", api.Msg.Pull)
		group.POST("/send", api.Msg.Send)
		group.GET("/chat", ws.ChatHandler)
	})
}
