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
	if entity.IsABotCommand(submittedMessageEvent.Message) {
		m.HandleBotCommand(submittedMessageEvent)
	} else {
		m.HandleMessage(submittedMessageEvent)
	}
}

func (m *MessageSender) HandleBotCommand(e *event.MessageSubmitted) {
	//TODO handle err
	submittedBotMessage, _ := entity.NewBotMessage(e.Message)
	sentBotMessage := event.NewBotMessageSent(submittedBotMessage)
	m.mediator.Publish(sentBotMessage)
}

func (m *MessageSender) HandleMessage(e *event.MessageSubmitted) {
	//TODO handle err
	submittedMessage, _ := entity.NewChatMessage(e.UserName, e.Message, e.Time)
	m.repo.CreateMessage(submittedMessage)
	sentMessage := event.NewMessageSent(submittedMessage.UserName, submittedMessage.Text, submittedMessage.Time)
	m.mediator.Publish(sentMessage)
}
