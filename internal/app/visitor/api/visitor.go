package api

import (
	"github.com/antbiz/antchat/internal/app/visitor/dto"
	"github.com/antbiz/antchat/internal/app/visitor/service"
	"github.com/antbiz/antchat/internal/app/ws"
	"github.com/antbiz/antchat/internal/db"
	"github.com/antbiz/antchat/internal/pkg/captcha"
	"github.com/antbiz/antchat/internal/pkg/resp"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Visitor = new(visitorApi)

type visitorApi struct{}

// Signin 访客登录
func (visitorApi) Signin(r *ghttp.Request) {
	var req *dto.VisitorSigninReq
	if err := r.Parse(&req); err != nil {
		resp.InvalidArgument(r, err.Error())
	}
	if !captcha.Verify(req.CaptchaID, req.Captcha) {
		resp.InvalidArgument(r, "验证码错误")
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
		visitor.ID = storeVisitor.ID
		visitor.Nickname = storeVisitor.Nickname
	} else {
		visitor.ID = primitive.NewObjectID()
	}

	selectOneAgent, err := service.SelectAgentID(ctx, visitor.AgentID)
	if err != nil {
		g.Log().Async().Errorf("visitor.Signin.SelectAgentID: %v", err)
	}
	visitor.AgentID = selectOneAgent

	if selectOneAgent != "" {
		// 通知客服
		ch := ws.AgentChatSrv().GetChannelByUID(visitor.AgentID)
		if ch != nil {
			err = ch.WriteSystemMessagef(visitor.AgentID, req.VisitorID, "来自 %s 的客户进入对话", visitor.Address())
			if err != nil {
				g.Log().Async().Errorf("visitor.Signin.NoticeAgentVisitorOnline: %v", err)
			}
		}
	}
	visitorID, err := db.UpsertVisitor(r.Context(), visitor)
	if err != nil {
		resp.DatabaseError(r, "保存信息失败")
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
