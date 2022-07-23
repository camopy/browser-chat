package messagereceiver

import (
	service "github.com/camopy/browser-chat/app/application/service"
	"github.com/camopy/browser-chat/app/domain/entity"
	"github.com/camopy/browser-chat/app/domain/event"
)

const eventName = "MessageSent"

type MessageReceiver struct {
	eventName   string
	broadcaster service.Broadcaster
}

func New(mediator service.Mediator, broadcaster service.Broadcaster) *MessageReceiver {
	return &MessageReceiver{
		eventName:   eventName,
		broadcaster: broadcaster,
	}
}

func (m *MessageReceiver) Name() string {
	return m.eventName
}

func (m *MessageReceiver) Handle(e event.DomainEvent) {
	sentMessageEvent := e.(*event.MessageSent)
	//TODO handle err
	sentMessage, _ := entity.NewChatMessage(sentMessageEvent.UserName, sentMessageEvent.Message, sentMessageEvent.Time)
	go m.broadcaster.Broadcast(sentMessage)
}
