package ws

import (
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Channel struct {
	uid  string
	send chan []byte
	conn *ghttp.WebSocket
	sess *ghttp.Session
}

func (ch *Channel) WriteMessage(msg *ChatMsg) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return ch.conn.WriteMessage(ghttp.WS_MSG_TEXT, body)
}

func (ch *Channel) WriteTextMessage(avatar, text string) error {
	msg := NewChatMsg("text", avatar, g.Map{
		"text": text,
	})
	return ch.WriteMessage(msg)
}

func (ch *Channel) WriteTextMessagef(avatar, format string, v ...interface{}) error {
	msg := NewChatMsg("text", avatar, g.Map{
		"text": fmt.Sprintf(format, v...),
	})
	return ch.WriteMessage(msg)
}

func (ch *Channel) WriteSystemMessage(text string) error {
	msg := NewChatMsg("system", "", g.Map{
		"text": text,
	})
	return ch.WriteMessage(msg)
}

func (ch *Channel) WriteSystemMessagef(text string, format string, v ...interface{}) error {
	msg := NewChatMsg("system", "", g.Map{
		"text": fmt.Sprintf(format, v...),
	})
	return ch.WriteMessage(msg)
}
