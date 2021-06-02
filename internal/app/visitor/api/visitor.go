package api

import (
	"github.com/antbiz/antchat/internal/app/visitor/dto"
	"github.com/antbiz/antchat/internal/app/visitor/service"
	"github.com/antbiz/antchat/internal/app/ws"
	"github.com/antbiz/antchat/internal/db"
	"github.com/antbiz/antchat/internal/pkg/resp"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Visitor = new(visitorApi)

type visitorApi struct{}

// Login 访客登录
func (visitorApi) Login(r *ghttp.Request) {
	var req *dto.VisitorLoginReq
	if err := r.Parse(&req); err != nil {
		resp.InvalidArgument(r, err.Error())
	}

	ctx := r.Context()
	visitor := &db.Visitor{
		Domain: req.Domain,
		Geo:    req.Geo,
		Refer:  r.GetReferer(),
		IP:     r.GetClientIp(),
	}
	// TODO: 解析ip地址信息
	visitor.Nickname = visitor.IP

	if req.VisitorID != "" {
		storeVisitor, err := db.GetVisitorByID(ctx, req.VisitorID)
		if err != nil {
			resp.DatabaseError(r, "查询访客失败")
		}
		visitor.AgentID = storeVisitor.AgentID
	}

	selectOneAgent, err := service.SelectAgentID(ctx, visitor.AgentID)
	if err != nil {
		g.Log().Async().Errorf("visitor.Login.SelectAgentID: %v", err)
	}
	visitor.AgentID = selectOneAgent

	visitorID, err := db.UpsertVisitor(r.Context(), req.VisitorID, visitor)
	if err != nil {
		resp.DatabaseError(r, "保存信息失败")
	}

	// 通知客服
	ch := ws.GetChannelByUID(visitor.AgentID)
	if ch != nil {
		err = ch.WriteSystemMessagef("来自 %s 的客户进入对话", visitor.Address())
		if err != nil {
			g.Log().Async().Errorf("visitor.Login.NoticeAgentVisitorOnline: %v", err)
		}
	}

	sessionData := g.Map{
		"id":        visitorID,
		"nickname":  visitor.Nickname,
		"sid":       r.Session.Id(),
		"agentID":   visitor.AgentID,
		"isVisitor": true,
	}
	r.Session.SetMap(sessionData)
	resp.OK(r, sessionData)
}
