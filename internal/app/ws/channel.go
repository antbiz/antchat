package ws

import (
	"github.com/gogf/gf/net/ghttp"
)

type Channel struct {
	uid  string
	send chan *Msg
	conn *ghttp.WebSocket
}

func (ch *Channel) WriteMessage(msg []byte) error {
	return ch.conn.WriteMessage(ghttp.WS_MSG_TEXT, msg)
}
