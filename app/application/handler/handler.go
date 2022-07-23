package handler

import "github.com/camopy/browser-chat/app/domain/event"

type Handler interface {
	Name() string
	Handle(e event.DomainEvent)
}
