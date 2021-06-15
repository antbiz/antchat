package api

import (
	"context"

	"github.com/antbiz/antchat/internal/app/visitor/dto"
	"github.com/antbiz/antchat/internal/app/ws"
	"github.com/antbiz/antchat/internal/db"
	"github.com/antbiz/antchat/internal/pkg/resp"
	"github.com/antbiz/antchat/internal/shared"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

var Msg = new(msgApi)

type msgApi struct{}

// Send 发送消息
func (msgApi) Send(r *ghttp.Request) {
	var req *ws.ChatMsg
	if err := r.Parse(&req); err != nil {
		resp.InvalidArgument(r, err.Error())
	}
	req.CreatedAt = gtime.Now().Timestamp()

	ctx := r.Context()
	ctxVisitor := shared.Ctx.GetCtxVisitor(ctx)
	// FIXME: 用ctx会报错：db.CreateMessage: connection(localhost:27017[-4]) failed to write: context canceled
	go db.CreateMessage(context.Background(), &db.Message{
		AgentID:   ctxVisitor.AgentID,
		VisitorID: ctxVisitor.ID,
		SenderID:  ctxVisitor.ID,
		Content:   req.Content,
		Type:      req.Type,
	})
	req.AgentID = ctxVisitor.AgentID
	req.VisitorID = ctxVisitor.ID

	g.Log().Async().Debugf("获取访客 %s session信息中的客服id：%s", ctxVisitor.ID, ctxVisitor.AgentID)
	ach := ws.AgentChatSrv().GetChannelByUID(ctxVisitor.AgentID)
	// TODO: 未被接待的访客放到队列等待分配
	if ach != nil {
		if err := ach.WriteMessage(req); err != nil {
			resp.InternalServer(r, "err_ws_write_msg", "发送失败")
		}
	} else {
		vch := ws.VisitorChatSrv().GetChannelByUID(ctxVisitor.ID)
		if vch != nil {
			_ = vch.WriteSystemMessage(ctxVisitor.AgentID, ctxVisitor.ID, "当前客服繁忙，请您耐心等待！")
		}
	}
	resp.OK(r)
}

// History 拉取消息列表
func (msgApi) History(r *ghttp.Request) {
	var req *dto.PullMsgReq
	if err := r.Parse(&req); err != nil {
		resp.InvalidArgument(r, err.Error())
	}

	ctx := r.Context()
	ctxVisitor := shared.Ctx.GetCtxVisitor(ctx)

	msgs, err := db.FindMessageByVisitorID(ctx, ctxVisitor.ID)
	if err != nil {
		resp.DatabaseError(r, "拉取消息失败")
	}
	resp.PageOK(r, 0, msgs)
}
