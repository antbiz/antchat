package app

import (
	"github.com/antbiz/antchat/internal/app/agent"
	"github.com/antbiz/antchat/internal/app/visitor"
	"github.com/antbiz/antchat/internal/middleware"
	"github.com/gogf/gf/frame/g"
)

func Run() {
	s := g.Server()
	s.Use(middleware.CORS)

	agent.Init()
	visitor.Init()

	s.Run()
}
