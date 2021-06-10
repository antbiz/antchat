package ws

import (
	"runtime"
	"time"

	"github.com/antbiz/antchat/internal/pkg/cityhash"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Server struct {
	Buckets    []*Bucket
	bucketIdx  uint32
	WriteWait  time.Duration
	PongWait   time.Duration
	PingPeriod time.Duration
	MaxMsgSize int64
	BufSize    int
}

// NewServer .
func NewServer() *Server {
	srv := &Server{
		WriteWait:  10 * time.Second,
		PongWait:   60 * time.Second,
		PingPeriod: (60 * time.Second * 9) / 10,
		MaxMsgSize: 512,
		BufSize:    256,
	}
	srv.Buckets = make([]*Bucket, g.Cfg().GetInt("ws.bucketSize", runtime.NumCPU()))
	srv.bucketIdx = uint32(len(srv.Buckets))

	bucketChannelSize := g.Cfg().GetInt("ws.bucketChannelSize", 1024)
	for i := 0; i < len(srv.Buckets); i++ {
		bucket := new(Bucket)
		bucket.chs = make(map[string]*Channel, bucketChannelSize)
		srv.Buckets[i] = bucket
	}
	return srv
}

func (srv *Server) Bucket(uid string) *Bucket {
	idx := cityhash.CityHash32([]byte(uid), uint32(len(uid))) % srv.bucketIdx
	g.Log().Async().Debugf("%s his channel bucket index: %d use cityhash", uid, idx)
	return srv.Buckets[idx]
}

func (srv *Server) GetChannelByUID(uid string) *Channel {
	if uid == "" {
		return nil
	}
	b := srv.Bucket(uid)
	b.cLock.RLock()
	ch := b.chs[uid]
	b.cLock.RUnlock()
	return ch
}

// readPump get data from websocket conn
func (srv *Server) readPump(ch *Channel) {
	defer func() {
		if ch.uid != "" {
			bucket := srv.Bucket(ch.uid)
			if bucket != nil {
				bucket.Del(ch.uid)
			} else {
				g.Log().Errorf("not found bucket for user %s", ch.uid)
			}
		}
		_ = ch.conn.Close()
		// 如果访客离开对话需要通知客服
		if ch.sess.GetBool("isVisitor") {
			g.Log().Async().Debugf("访客 %s 关闭ws连接, 通知客服 %s", ch.uid, ch.sess.GetString("agentID"))
			agentCh := agentChatSrv.GetChannelByUID(ch.sess.GetString("agentID"))
			if agentCh != nil {
				if err := agentCh.WriteSystemMessagef("客户 %s 关闭对话", ch.sess.GetString("nickname")); err != nil {
					g.Log().Async().Errorf("通知客服 %s 访客 %s 关闭对话：%v", ch.sess.GetString("agentID"), ch.uid, err)
				}
			} else {
				g.Log().Debugf("无法获取客服 %s 的回话", ch.sess.GetString("agentID"))
			}

		} else {
			g.Log().Async().Debugf("客服 %s 关闭ws连接", ch.uid)
		}
	}()

	ch.conn.SetReadLimit(srv.MaxMsgSize)
	ch.conn.SetReadDeadline(time.Now().Add(srv.PongWait))
	ch.conn.SetPongHandler(func(string) error {
		ch.conn.SetReadDeadline(time.Now().Add(srv.PongWait))
		return nil
	})

	for {
		_, message, err := ch.conn.ReadMessage()
		if err != nil {
			return
		}
		if message == nil {
			return
		}
	}
}

// writePump send data to websocket conn
func (srv *Server) writePump(ch *Channel) {
	ticker := time.NewTicker(srv.PingPeriod)
	defer func() {
		ticker.Stop()
		_ = ch.conn.Close()
	}()

	for {
		select {
		case message, ok := <-ch.send:
			ch.conn.SetWriteDeadline(time.Now().Add(srv.WriteWait))
			if !ok {
				ch.conn.WriteMessage(ghttp.WS_MSG_CLOSE, []byte{})
				return
			}

			w, err := ch.conn.NextWriter(ghttp.WS_MSG_TEXT)
			if err != nil {
				g.Log().Async().Errorf("ch.conn.NextWriter: %v", err)
				return
			}
			w.Write(message)

			n := len(ch.send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-ch.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			g.Log().Async().Debug("websocket ping....")
			ch.conn.SetWriteDeadline(time.Now().Add(srv.WriteWait))
			if err := ch.conn.WriteMessage(ghttp.WS_MSG_PING, nil); err != nil {
				return
			}
		}
	}
}
