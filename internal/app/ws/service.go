package ws

import (
	"context"
	"time"

	"github.com/antbiz/antchat/internal/db"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

type Conversation struct {
	VisitorID string                 `json:"id"`
	Nickname  string                 `json:"nickname"`
	Content   map[string]interface{} `json:"content"`
	ActiveAt  time.Time              `json:"activeAt"`
}

// GetRealtimeConversations 获取当前所有对话
func GetRealtimeConversations(ctx context.Context) ([]*Conversation, error) {
	// 获取当前在线的所有的访客
	onlineVisitorIDs := make([]string, 0)
	for _, b := range visitorChatSrv.Buckets {
		for _, visitor := range b.chs {
			onlineVisitorIDs = append(onlineVisitorIDs, visitor.uid)
		}
	}
	if len(onlineVisitorIDs) == 0 {
		return nil, nil
	}
	visitorNicks, err := db.GetVisitorNicks(ctx, onlineVisitorIDs)
	if err != nil {
		return nil, err
	}
	msgs, err := db.GetLastMessagesByVisitorIDs(ctx, onlineVisitorIDs)
	if err != nil {
		return nil, err
	}
	visitorMsgs := make(map[string]*db.Message, len(msgs))
	for _, msg := range msgs {
		visitorMsgs[msg.VisitorID] = msg
	}

	// TODO: 按照加入时间和会话时间排序

	conversations := make([]*Conversation, len(onlineVisitorIDs))
	for i, vid := range onlineVisitorIDs {
		msg := visitorMsgs[vid]
		if msg == nil {
			conversations[i] = &Conversation{
				VisitorID: vid,
				Nickname:  visitorNicks[vid],
				Content: g.Map{
					"text": "新访客",
				},
				ActiveAt: gtime.Now().Time,
			}
		} else {
			conversations[i] = &Conversation{
				VisitorID: msg.VisitorID,
				Nickname:  visitorNicks[msg.VisitorID],
				Content:   msg.Content,
				ActiveAt:  msg.CreatedAt,
			}
		}
	}

	return conversations, nil
}
