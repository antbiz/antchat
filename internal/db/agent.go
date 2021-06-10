package db

import (
	"context"

	"github.com/gogf/gf/frame/g"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Agent 客服
type Agent struct {
	DefaultField `bson:",inline"`
	UserID       string `bson:"userID" json:"userID"`
	Status       int    `bson:"status" json:"status"`
	Online       bool   `bson:"online" json:"online"`
	Blocked      bool   `bson:"blocked" json:"blocked"`
}

func GetAgentCollection() *qmgo.Collection {
	return DB().Collection("agent")
}

func CreateAgent(ctx context.Context, msg *Agent) (id string, err error) {
	res, err := GetAgentCollection().InsertOne(ctx, msg)
	if err != nil {
		g.Log().Errorf("db.CreateAgent: %v", err)
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func UpsertAgent(ctx context.Context, agentID string, agent *Agent) (id string, err error) {
	res, err := GetAgentCollection().UpsertId(ctx, agent.ID, agent)
	if err != nil {
		g.Log().Async().Errorf("db.UpsertAgent: %v", err)
		return "", err
	}
	if res.UpsertedID == nil {
		id = agentID
	} else {
		id = res.UpsertedID.(primitive.ObjectID).Hex()
	}
	return
}

func GetAgentByUID(ctx context.Context, uid string) (agent *Agent, err error) {
	err = GetAgentCollection().
		Find(
			ctx,
			bson.M{"userID": uid},
		).One(&agent)
	return
}

func GetAllAgents(ctx context.Context, filter interface{}) ([]*Agent, error) {
	agents := make([]*Agent, 0)
	err := GetAgentCollection().Find(ctx, filter).All(&agents)
	if err != nil {
		return nil, err
	}
	return agents, nil
}
