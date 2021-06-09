package service

import (
	"context"

	"github.com/antbiz/antchat/internal/db"
	"github.com/gogf/gf/util/grand"
	"go.mongodb.org/mongo-driver/bson"
)

// SelectAgentID 选择一个客服对接，优先选择在线的客户
func SelectAgentID(ctx context.Context, oldAgentID ...string) (string, error) {
	var oldID string
	if len(oldAgentID) > 0 {
		oldID = oldAgentID[0]
	}
	agents, err := db.GetAllAgents(ctx, bson.M{})
	if len(agents) == 0 || err != nil {
		return "", err
	}

	onlineAgentIDs := make([]string, 0)
	agentIDs := make([]string, len(agents))
	for i, agent := range agents {
		agentStrID := agent.ID.Hex()
		if agent.Online {
			if agentStrID == oldID {
				return oldID, nil
			}
			onlineAgentIDs = append(onlineAgentIDs, agentStrID)
		}
		agentIDs[i] = agentStrID
	}

	// TODO: 优化这里，目前是随机选择
	onlineCount := len(onlineAgentIDs)
	if onlineCount > 0 {
		return onlineAgentIDs[grand.Intn(onlineCount)], nil
	}
	return agentIDs[grand.Intn(onlineCount)], nil
}
