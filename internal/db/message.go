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
	AgentID      string                 `bson:"agentID"`    // 客服ID
	VisitorID    string                 `bson:"visitorID"`  // 访客ID
	SenderID     string                 `bson:"senderID"`   // 发送者ID
	SenderNick   string                 `bson:"senderNick"` // 发送者昵称
	Type         string                 `bson:"type"`       // 消息类型
	Content      map[string]interface{} `bson:"content"`    // 消息内容
	Status       int                    `bson:"status"`     // 消息状态
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

func GetLastMessagesByVisitorIDs(ctx context.Context, ids []string) ([]*Message, error) {
	sortState := bson.D{{"$sort", bson.M{"createdAt": -1}}}
	matchStage := bson.D{{"$match", []bson.E{{"visitorID", bson.D{{"$in", ids}}}}}}
	groupStage := bson.D{{"$group", bson.D{
		{"_id", "$visitorID"},
		{"doc", bson.M{"$first": "$$ROOT"}},
	}}}
	// https://stackoverflow.com/a/59756228
	replaceState := bson.D{{"$replaceRoot", bson.M{"newRoot": "$doc"}}}

	msgs := make([]*Message, 0)

	err := GetMessageCollection().
		Aggregate(
			ctx,
			qmgo.Pipeline{sortState, matchStage, groupStage, replaceState},
		).
		All(&msgs)
	return msgs, err
}
