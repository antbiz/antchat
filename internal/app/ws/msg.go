package ws

const (
	ChatMsgTypeSystem = "system"
	ChatMsgTypeText   = "text"
	ChatMsgTypeImage  = "image"
	ChatMsgTypeCmd    = "cmd"
)

type ChatMsg struct {
	Type      string       `json:"type"`
	Content   interface{}  `json:"content" v:"required"`
	CreatedAt int64        `json:"createdAt"`
	User      *ChatMsgUser `json:"user"`
}

type ChatMsgUser struct {
	Avatar string `json:"avatar"`
}
