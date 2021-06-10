package ws

import (
	"context"
	"time"

	"github.com/antbiz/antchat/internal/db"
	"github.com/gogf/gf/os/gtime"
)

type Conversation struct {
	VisitorID string      `json:"id"`
	Nickname  string      `json:"nickname"`
	Message   interface{} `json:"message"`
	ActiveAt  time.Time   `json:"activeAt"`
}

// GetRealtimeConversations 获取当前所有对话
func GetRealtimeConversations(ctx context.Context) ([]*Conversation, error) {
	// 获取当前在线的所有的访客
	onlineVisitorIDs := make([]string, 0)
	visitorNicks := make(map[string]string)
	for _, b := range visitorChatSrv.Buckets {
		for _, visitor := range b.chs {
			onlineVisitorIDs = append(onlineVisitorIDs, visitor.uid)
			visitorNicks[visitor.uid] = visitor.sess.GetString("nickname")
		}
	}
	if len(onlineVisitorIDs) == 0 {
		return nil, nil
	}

	msgs, err := db.GetLastMessagesByVisitorIDs(ctx, onlineVisitorIDs)
	if err != nil {
		return nil, err
	}
	visitorMsgs := make(map[string]*db.Message, len(onlineVisitorIDs))
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
				Message:   "",
				ActiveAt:  gtime.Now().Time,
			}
		} else {
			conversations[i] = &Conversation{
				VisitorID: msg.VisitorID,
				Nickname:  visitorNicks[msg.VisitorID],
				Message:   msg.Content,
				ActiveAt:  msg.CreatedAt,
			}
		}
	}

	return conversations, nil
}
