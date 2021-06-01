package ws

import (
	"time"

	"github.com/antbiz/antchat/internal/pkg/cityhash"
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

func (srv *Server) Bucket(uid string) *Bucket {
	idx := cityhash.CityHash32([]byte(uid), uint32(len(uid))) % srv.bucketIdx
	return srv.Buckets[idx]
}

// readPump get data from websocket conn
func (srv *Server) readPump(ch *Channel) {
	defer func() {
		if ch.uid != "" {
			srv.Bucket(ch.uid).Del(ch.uid)
		}
		_ = ch.conn.Close()
		// 如果访客离开对话需要通知客服
		if ch.sess.GetBool("isVisitor") {
			agentCh := GetChannelByUID(ch.sess.GetString("agentID"))
			if agentCh != nil {
				_ = agentCh.WriteSystemMessagef("客户 %s 关闭对话", ch.sess.GetString("nickname"))
			}
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
			}

			w, err := ch.conn.NextWriter(ghttp.WS_MSG_TEXT)
			if err != nil {
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
			ch.conn.SetWriteDeadline(time.Now().Add(srv.WriteWait))
			if err := ch.conn.WriteMessage(ghttp.WS_MSG_PING, nil); err != nil {
				return
			}
		}
	}
}
