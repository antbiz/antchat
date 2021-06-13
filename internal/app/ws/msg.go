package ws

import "github.com/gogf/gf/os/gtime"

const (
	ChatMsgTypeSystem = "system"
	ChatMsgTypeText   = "text"
	ChatMsgTypeImage  = "image"
	ChatMsgTypeCmd    = "cmd"
)

type ChatMsg struct {
	AgentID   string                 `json:"agentID"`
	VisitorID string                 `json:"visitorID"`
	Type      string                 `json:"type"`
	Content   map[string]interface{} `json:"content" v:"required"`
	CreatedAt int64                  `json:"createdAt"`
}

func NewChatMsg(aid, vid, msgType string, content map[string]interface{}) *ChatMsg {
	msg := &ChatMsg{
		AgentID:   aid,
		VisitorID: vid,
		Type:      msgType,
		Content:   content,
		CreatedAt: gtime.Now().Timestamp(),
	}
	return msg
}
