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
	g.Log().Async().Debug("初始化访客websocket服务")
	visitorChatSrv = NewServer()
	g.Log().Async().Debug("初始化客服websocket服务")
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
		send: make(chan []byte, visitorChatSrv.BufSize),
	}
	b := visitorChatSrv.Bucket(ch.uid)
	b.Set(ch.uid, ch)

	// 新访客加入，通知客服
	aid := ch.sess.GetString("agentID")
	g.Log().Async().Debugf("通知客服 %s, 访客 %s 已连接", aid, ch.uid)
	agentCh := agentChatSrv.GetChannelByUID(ch.sess.GetString("agentID"))
	if agentCh != nil {
		var (
			lastMsgContent map[string]interface{}
			activeAt       time.Time
			visitorNick    string
		)
		lastMsg, _ := db.GetLastMessageByVisitorID(ctx, ctxVisitor.ID)
		if lastMsg == nil {
			lastMsgContent = nil
			activeAt = gtime.Now().Time
		} else {
			lastMsgContent = lastMsg.Content
			activeAt = lastMsg.CreatedAt
		}
		if visitor, err := db.GetVisitorByID(ctx, ctxVisitor.ID); err != nil {
			g.Log().Async().Errorf("ws.VisitorChatHandler.GetVisitorByID: %v", err)
		} else {
			visitorNick = visitor.Nickname
		}

		msg := NewChatMsg(aid, ch.uid, ChatMsgTypeCmd, g.Map{
			"data": &Conversation{
				VisitorID: ctxVisitor.ID,
				Nickname:  visitorNick,
				Content:   lastMsgContent,
				ActiveAt:  activeAt,
			},
			"code": "incoming_update",
		})
		if err = agentCh.WriteMessage(msg); err != nil {
			g.Log().Async().Errorf("通知客服 %s 失败: %v", ch.sess.GetString("agentID"), err)
		}
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

	// NOTE: 客服的uid使用客服会话中的客服id，和访客保持一致
	ch := &Channel{
		uid:  r.Session.GetString("agentID"),
		conn: ws,
		sess: r.Session,
		send: make(chan []byte, agentChatSrv.BufSize),
	}
	b := agentChatSrv.Bucket(ch.uid)
	b.Set(ch.uid, ch)
	g.Log().Async().Debugf("客服 %s 已连接", ch.uid)

	go agentChatSrv.writePump(ch)
	go agentChatSrv.readPump(ch)
}
