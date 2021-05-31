package api

import (
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
	var req *dto.SendMsgRequest
	if err := r.Parse(&req); err != nil {
		resp.InvalidArgument(r, err.Error())
	}

	ctx := r.Context()
	ctxUser := shared.Ctx.Get(ctx).User

	req.SenderID = ctxUser.ID
	req.SenderNick = ctxUser.Nickname
	go db.CreateMessage(ctx, &db.Message{
		AgentID:    req.SenderID,
		VisitorID:  req.ReceiverID,
		SenderID:   req.SenderID,
		SenderRole: req.SenderRole,
		SenderNick: req.SenderNick,
		Content:    req.Body,
	})

	ch := ws.GetChannelByUID(req.ReceiverID)
	if ch == nil {
		resp.OK(r, "对方已断开")
	}
	if err := ch.WriteMessage(req.Body); err != nil {
		resp.InternalServer(r, "err_ws_write_msg", "发送失败")
	}
	resp.OK(r)
}

// Pull 拉取消息列表
func (msgApi) Pull(r *ghttp.Request) {
	var req *dto.PullMsgRequest
	if err := r.Parse(&req); err != nil {
		resp.InvalidArgument(r, err.Error())
	}

	ctx := r.Context()

	msgs, err := db.FindMessageByVisitorID(ctx, req.VisitorID)
	if err != nil {
		resp.DatabaseError(r, "拉取消息失败")
	}
	resp.PageOK(r, 0, msgs)
}
