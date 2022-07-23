package mediator

import (
	"github.com/camopy/browser-chat/app/application/handler"
	"github.com/camopy/browser-chat/app/domain/event"
)

type Mediator struct {
	handlers []handler.Handler
}

func New() *Mediator {
	return &Mediator{
		handlers: []handler.Handler{},
	}
}

func (m *Mediator) Register(h handler.Handler) {
	m.handlers = append(m.handlers, h)
}

func (m *Mediator) Publish(e event.DomainEvent) {
	for _, h := range m.handlers {
		if h.Name() == e.Name() {
			h.Handle(e)
		}
	}
}
