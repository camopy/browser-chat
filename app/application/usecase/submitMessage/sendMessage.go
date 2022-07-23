package submitmessage

import (
	"time"

	mediator "github.com/camopy/browser-chat/app/application/service"
	"github.com/camopy/browser-chat/app/domain/event"
)

type SubmitMessage struct {
	mediator mediator.Mediator
}

type Input struct {
	UserName string
	Message  string
	Time     time.Time
}

func New(m mediator.Mediator) *SubmitMessage {
	return &SubmitMessage{
		mediator: m,
	}
}

func (m *SubmitMessage) Execute(input Input) {
	event := event.NewMessageSubmitted(input.UserName, input.Message, input.Time)
	m.mediator.Publish(event)
}
