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
		g.Log().Async().Errorf("ws.channel.WriteMessage.json.Marshal: %v", err)
		return err
	}
	ch.send <- body
	return nil
}

func (ch *Channel) WriteTextMessage(aid, vid, text string) error {
	msg := NewChatMsg(aid, vid, ChatMsgTypeText, g.Map{
		"text": text,
	})
	return ch.WriteMessage(msg)
}

func (ch *Channel) WriteTextMessagef(aid, vid, format string, v ...interface{}) error {
	msg := NewChatMsg(aid, vid, ChatMsgTypeText, g.Map{
		"text": fmt.Sprintf(format, v...),
	})
	return ch.WriteMessage(msg)
}

func (ch *Channel) WriteSystemMessage(aid, vid, text string) error {
	msg := NewChatMsg(aid, vid, ChatMsgTypeSystem, g.Map{
		"text": text,
	})
	return ch.WriteMessage(msg)
}

func (ch *Channel) WriteSystemMessagef(aid, vid, text string, format string, v ...interface{}) error {
	msg := NewChatMsg(aid, vid, ChatMsgTypeSystem, g.Map{
		"text": fmt.Sprintf(format, v...),
	})
	return ch.WriteMessage(msg)
}
