package main

import (
	_ "antchat/boot"
	_ "antchat/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
