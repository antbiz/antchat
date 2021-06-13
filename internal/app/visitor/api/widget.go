package api

import "github.com/gogf/gf/net/ghttp"

var Widget = new(widgetApi)

type widgetApi struct{}

func (widgetApi) Index(r *ghttp.Request) {
	r.Response.WriteTpl("widget.html")
}
