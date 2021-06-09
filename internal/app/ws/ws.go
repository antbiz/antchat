package ws

import (
	"time"

	"github.com/antbiz/antchat/internal/db"
	"github.com/antbiz/antchat/internal/shared"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

var (
	visitorChatSrv *Server
	agentChatSrv   *Server
)

func init() {
	visitorChatSrv = NewServer()
	agentChatSrv = NewServer()
}

// VisitorChatSrv .
func VisitorChatSrv() *Server {
	return visitorChatSrv
}

// AgentChatSrv .
func AgentChatSrv() *Server {
	return agentChatSrv
}

func VisitorChatHandler(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		g.Log().Error(err)
		return
	}

	ctx := r.Context()
	ctxVisitor := shared.Ctx.GetCtxVisitor(ctx)

	ch := &Channel{
		uid:  ctxVisitor.ID,
		conn: ws,
		sess: r.Session,
	}
	b := visitorChatSrv.Bucket(ch.uid)
	b.Set(ch.uid, ch)

	// 新访客加入，通知客服
	agentCh := agentChatSrv.GetChannelByUID(ch.sess.GetString("agentID"))
	if agentCh != nil {
		var (
			lastMsgContent interface{}
			activeAt       time.Time
		)
		lastMsg, _ := db.GetLastMessageByVisitorID(ctx, ctxVisitor.ID)
		if lastMsg == nil {
			lastMsgContent = ""
			activeAt = gtime.Now().Time
		} else {
			lastMsgContent = lastMsg.Content
			activeAt = lastMsg.CreatedAt
		}

		msg := NewChatMsg(ChatMsgTypeSystem, "", g.Map{
			"data": g.Map{
				"id":       ctxVisitor.ID,
				"nickname": ctxVisitor.Nickname,
				"message":  lastMsgContent,
				"activeAt": activeAt,
			},
			"code": "incoming_update",
		})
		_ = agentCh.WriteMessage(msg)
	}

	go visitorChatSrv.writePump(ch)
	go visitorChatSrv.readPump(ch)
}

func AgentChatHandler(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		g.Log().Error(err)
		return
	}

	ch := &Channel{
		uid:  r.Session.GetString("id"),
		conn: ws,
		sess: r.Session,
	}
	b := agentChatSrv.Bucket(ch.uid)
	b.Set(ch.uid, ch)

	go agentChatSrv.writePump(ch)
	go agentChatSrv.readPump(ch)
}
