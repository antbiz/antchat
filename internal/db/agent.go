package db

import (
	"context"

	"github.com/qiniu/qmgo"
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

func GetAllAgents(ctx context.Context, filter interface{}) ([]*Agent, error) {
	agents := make([]*Agent, 0)
	err := GetAgentCollection().Find(ctx, filter).All(&agents)
	if err != nil {
		return nil, err
	}
	return agents, nil
}
