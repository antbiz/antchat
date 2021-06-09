package main

import (
	"github.com/antbiz/antchat/internal/app"
	"github.com/antbiz/antchat/internal/db"
	"github.com/gogf/gf/frame/g"
)

func main() {
	if err := db.Cli().Ping(10); err != nil {
		g.Log().Fatalf("ping mongo failed: %v", err)
	}
	app.Run()
}
