package ws

import (
	"sync"
)

type Bucket struct {
	cLock       sync.RWMutex
	chs         map[string]*Channel
	ChannelSize int
}

func (b *Bucket) Set(uid string, ch *Channel) {
	b.cLock.Lock()
	ch.uid = uid
	b.chs[uid] = ch
	b.cLock.Unlock()
}

func (b *Bucket) Del(uid string) {
	b.cLock.Lock()
	delete(b.chs, uid)
	b.cLock.Unlock()
}

func GetChannelByUID(uid string) *Channel {
	b := chatSrv.Bucket(uid)
	b.cLock.RLock()
	ch := b.chs[uid]
	b.cLock.RUnlock()
	return ch
}
