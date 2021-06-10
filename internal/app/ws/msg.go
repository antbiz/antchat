package ws

import "github.com/gogf/gf/os/gtime"

const (
	ChatMsgTypeSystem = "system"
	ChatMsgTypeText   = "text"
	ChatMsgTypeImage  = "image"
	ChatMsgTypeCmd    = "cmd"
)

type ChatMsg struct {
	Type      string                 `json:"type"`
	Content   map[string]interface{} `json:"content" v:"required"`
	CreatedAt int64                  `json:"createdAt"`
	User      struct {
		Avatar string `json:"avatar"`
	} `json:"user"`
}

func NewChatMsg(msgType, avatar string, content map[string]interface{}) *ChatMsg {
	msg := &ChatMsg{
		Type:      msgType,
		Content:   content,
		CreatedAt: gtime.Now().Timestamp(),
	}
	msg.User.Avatar = avatar
	return msg
}
