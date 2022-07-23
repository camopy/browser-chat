package messagesender

import (
	mediator "github.com/camopy/browser-chat/app/application/service"
	"github.com/camopy/browser-chat/app/domain/entity"
	"github.com/camopy/browser-chat/app/domain/event"
	"github.com/camopy/browser-chat/app/domain/repository"
)

const eventName = "MessageSubmitted"

type MessageSender struct {
	eventName string
	mediator  mediator.Mediator
	repo      repository.ChatMessageRepository
}

func New(mediator mediator.Mediator, repo repository.ChatMessageRepository) *MessageSender {
	return &MessageSender{
		eventName: eventName,
		mediator:  mediator,
		repo:      repo,
	}
}

func (m *MessageSender) Name() string {
	return m.eventName
}

func (m *MessageSender) Handle(e event.DomainEvent) {
	submittedMessageEvent := e.(*event.MessageSubmitted)
	submittedMessage, _ := entity.NewChatMessage(submittedMessageEvent.UserName, submittedMessageEvent.Message, submittedMessageEvent.Time)
	m.repo.Save(submittedMessage)
}
