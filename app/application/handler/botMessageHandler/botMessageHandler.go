package botmessagehandler

import (
	"github.com/camopy/browser-chat/app/application/bot"
	"github.com/camopy/browser-chat/app/application/service"
	"github.com/camopy/browser-chat/app/domain/event"
)

const eventName = "BotMessageSent"

type BotMessageHandler struct {
	eventName string
	mediator  service.Mediator
	bot       bot.Bot
}

func New(mediator service.Mediator, bot bot.Bot) *BotMessageHandler {
	return &BotMessageHandler{
		eventName: eventName,
		mediator:  mediator,
		bot:       bot,
	}
}

func (m *BotMessageHandler) Name() string {
	return m.eventName
}

func (m *BotMessageHandler) Handle(e event.DomainEvent) {
	sentBotMessageEvent := e.(*event.BotMessageSent)
	sentBotMessage := sentBotMessageEvent.Message
	//TODO handle err
	close, _ := m.bot.GetQuote(sentBotMessage.StockCode)
	sentBotMessage.Answer(close)
	sentMessage := event.NewMessageSent(sentBotMessage.BotName, sentBotMessage.Text, sentBotMessage.Time)
	m.mediator.Publish(sentMessage)
}
