package ws

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

type Channel struct {
	uid  string
	send chan []byte
	conn *ghttp.WebSocket
}

func (ch *Channel) WriteMessage(msg *ChatMsg) error {
	return ch.conn.WriteMessage(ghttp.WS_MSG_TEXT, gconv.Bytes(msg))
}
