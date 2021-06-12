package api

import (
	"context"

	"github.com/antbiz/antchat/internal/app/agent/dto"
	"github.com/antbiz/antchat/internal/app/ws"
	"github.com/antbiz/antchat/internal/db"
	"github.com/antbiz/antchat/internal/pkg/resp"
	"github.com/antbiz/antchat/internal/shared"
	"github.com/gogf/gf/net/ghttp"
)

var Msg = new(msgApi)

type msgApi struct{}

// Send 发送消息
func (msgApi) Send(r *ghttp.Request) {
	var req *dto.SendMsgReq
	if err := r.Parse(&req); err != nil {
		resp.InvalidArgument(r, err.Error())
	}

	ctx := r.Context()
	ctxUser := shared.Ctx.GetCtxUser(ctx)

	req.SenderID = ctxUser.ID
	req.SenderNick = ctxUser.Nickname
	go db.CreateMessage(context.Background(), &db.Message{
		AgentID:    ctxUser.AgentID,
		VisitorID:  req.ReceiverID,
		SenderID:   req.SenderID,
		SenderNick: req.SenderNick,
		Content:    req.Content,
		Type:       req.Type,
	})

	ch := ws.VisitorChatSrv().GetChannelByUID(req.ReceiverID)
	if ch == nil {
		resp.OK(r, "访客已关闭对话")
	}

	if err := ch.WriteMessage(ws.NewChatMsg(ctxUser.AgentID, req.ReceiverID, ctxUser.Avatar, req.Type, req.Content)); err != nil {
		resp.InternalServer(r, "err_ws_write_msg", "发送失败")
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

	// TODO: 前端实现历史消息下拉
	// msgs, err := db.FindMessageByVisitorIDWithPaging(ctx, req.VisitorID, req.PageNum, req.PageSize)
	msgs, err := db.FindMessageByVisitorID(ctx, req.VisitorID)
	if err != nil {
		resp.DatabaseError(r, "拉取消息失败")
	}
	resp.PageOK(r, len(msgs), msgs)
}

// Conversations 拉取对话列表
func (msgApi) Conversations(r *ghttp.Request) {
	// TODO: 仅拉取当前客服的对话列表
	res, err := ws.GetRealtimeConversations(r.Context())
	if err != nil {
		resp.InternalServer(r, "err_get_conversations", "拉取对话列表失败", err)
	}
	resp.PageOK(r, len(res), res)
}
