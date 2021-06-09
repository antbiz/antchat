package db

import (
	"context"

	"github.com/gogf/gf/frame/g"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

// Message 消息
type Message struct {
	DefaultField `bson:",inline"`
	AgentID      string      `bson:"agentID"`    // 客服ID
	VisitorID    string      `bson:"visitorID"`  // 访客ID
	SenderID     string      `bson:"senderID"`   // 发送者ID
	SenderRole   int         `bson:"senderRole"` // 发送者类型
	SenderNick   string      `bson:"senderNick"` // 发送者昵称
	Type         string      `bson:"type"`       // 消息类型
	Content      interface{} `bson:"content"`    // 消息内容
	Status       int         `bson:"status"`     // 消息状态
}

func GetMessageCollection() *qmgo.Collection {
	return DB().Collection("message")
}

func CreateMessage(ctx context.Context, msg *Message) {
	if _, err := GetMessageCollection().InsertOne(ctx, msg); err != nil {
		g.Log().Errorf("db.CreateMessage: %v", err)
	}
}

func FindMessageByVisitorID(ctx context.Context, id string) ([]*Message, error) {
	var msgs []*Message
	err := GetMessageCollection().Find(ctx, bson.M{"visitorID": id}).All(&msgs)
	return msgs, err
}

func GetLastMessageByVisitorID(ctx context.Context, id string) (*Message, error) {
	var msg *Message
	err := GetMessageCollection().Find(ctx, bson.M{"visitorID": id}).Select(bson.M{"createdAt": -1}).One(&msg)
	return msg, err
}
