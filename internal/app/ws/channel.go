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
	select {
	case ch.send <- body:
	default:
	}
	return nil
	// return ch.conn.WriteMessage(ghttp.WS_MSG_TEXT, body)
}

func (ch *Channel) WriteTextMessage(avatar, text string) error {
	msg := NewChatMsg(ChatMsgTypeText, avatar, g.Map{
		"text": text,
	})
	return ch.WriteMessage(msg)
}

func (ch *Channel) WriteTextMessagef(avatar, format string, v ...interface{}) error {
	msg := NewChatMsg(ChatMsgTypeText, avatar, g.Map{
		"text": fmt.Sprintf(format, v...),
	})
	return ch.WriteMessage(msg)
}

func (ch *Channel) WriteSystemMessage(text string) error {
	msg := NewChatMsg(ChatMsgTypeSystem, "", g.Map{
		"text": text,
	})
	return ch.WriteMessage(msg)
}

func (ch *Channel) WriteSystemMessagef(text string, format string, v ...interface{}) error {
	msg := NewChatMsg(ChatMsgTypeSystem, "", g.Map{
		"text": fmt.Sprintf(format, v...),
	})
	return ch.WriteMessage(msg)
}
