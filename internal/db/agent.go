package db

import (
	"context"

	"github.com/gogf/gf/container/garray"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

// Agent 客服
type Agent struct {
	DefaultField `bson:",inline"`
	UserID       string `bson:"userID"`
	Status       int    `bson:"status"`
	Online       bool   `bson:"online"`
	Blocked      bool   `bson:"blocked"`
}

func GetAgentCollection() *qmgo.Collection {
	return DB().Collection("agent")
}

func GetOnlineAgents(ctx context.Context) ([]*Agent, error) {
	agents := make([]*Agent, 0)
	err := GetAgentCollection().Find(ctx, bson.M{"online": true}).All(&agents)
	if err != nil {
		return nil, err
	}
	return agents, nil
}

func GetOnlineAgentIDs(ctx context.Context) (*garray.StrArray, error) {
	agents, err := GetOnlineAgents(ctx)
	if err != nil {
		return nil, err
	}
	ids := garray.NewStrArray()
	for _, agent := range agents {
		ids.Append(agent.ID)
	}
	return ids, nil
}
