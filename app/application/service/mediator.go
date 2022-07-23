package service

import (
	"github.com/camopy/browser-chat/app/application/handler"
	"github.com/camopy/browser-chat/app/domain/event"
)

type Mediator interface {
	Register(h handler.Handler)
	Publish(e event.DomainEvent)
}
