package event

import "github.com/camopy/browser-chat/app/domain/entity"

const botMessageSent = "BotMessageSent"

type BotMessageSent struct {
	name    string
	Message *entity.BotMessage
}

func NewBotMessageSent(message *entity.BotMessage) *BotMessageSent {
	return &BotMessageSent{
		name:    botMessageSent,
		Message: message,
	}
}

func (m *BotMessageSent) Name() string {
	return m.name
}
