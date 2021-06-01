package ws

import (
	"runtime"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var chatSrv *Server

func init() {
	chatSrv = &Server{
		WriteWait:  10 * time.Second,
		PongWait:   60 * time.Second,
		PingPeriod: (60 * 9) / 10,
		MaxMsgSize: 512,
		BufSize:    256,
	}
	chatSrv.Buckets = make([]*Bucket, g.Cfg().GetInt("ws.bucketSize", runtime.NumCPU()))
	chatSrv.bucketIdx = uint32(len(chatSrv.Buckets))

	bucketChannelSize := g.Cfg().GetInt("ws.bucketChannelSize", 1024)
	for i := 0; i < len(chatSrv.Buckets); i++ {
		bucket := new(Bucket)
		bucket.chs = make(map[string]*Channel, bucketChannelSize)
		chatSrv.Buckets[i] = bucket
	}
}

func ChatHandler(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		g.Log().Error(err)
		return
	}

	ch := &Channel{
		uid:  r.Session.GetString("id"),
		conn: ws,
	}
	b := chatSrv.Bucket(ch.uid)
	b.Set(ch.uid, ch)

	go chatSrv.writePump(ch)
	go chatSrv.readPump(ch)
}
