package api

import (
	"github.com/antbiz/antchat/internal/app/ws"
	"github.com/antbiz/antchat/internal/db"
	"github.com/antbiz/antchat/internal/pkg/resp"
	"github.com/antbiz/antchat/internal/shared"
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
	go db.CreateMessage(ctx, &db.Message{
		SenderID:   ctxVisitor.ID,
		SenderNick: ctxVisitor.Nickname,
	})

	ch := ws.GetChannelByUID(ctxVisitor.AgentID)
	if ch == nil {
		resp.OK(r, "对方已断开")
	}
	if err := ch.WriteMessage(req); err != nil {
		resp.InternalServer(r, "err_ws_write_msg", "发送失败")
	}
	resp.OK(r)
}

// Pull 拉取消息列表
func (msgApi) Pull(r *ghttp.Request) {
	ctx := r.Context()
	ctxVisitor := shared.Ctx.GetCtxVisitor(ctx)

	msgs, err := db.FindMessageByVisitorID(ctx, ctxVisitor.ID)
	if err != nil {
		resp.DatabaseError(r, "拉取消息失败")
	}
	resp.PageOK(r, 0, msgs)
}
