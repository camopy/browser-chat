package broadcaster

import (
	"github.com/camopy/browser-chat/app/domain/entity"
)

type Broadcaster struct {
	ch chan *entity.ChatMessage
}

func New() *Broadcaster {
	return &Broadcaster{
		ch: make(chan *entity.ChatMessage),
	}
}

func (b *Broadcaster) Broadcast(m *entity.ChatMessage) {
	b.ch <- m
}

func (b *Broadcaster) Receive() *entity.ChatMessage {
	return <-b.ch
}
