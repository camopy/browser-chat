package main

import (
	appbot "github.com/camopy/browser-chat/app/application/bot"
	botmessagehandler "github.com/camopy/browser-chat/app/application/handler/botMessageHandler"
	messagereceiver "github.com/camopy/browser-chat/app/application/handler/messageReceiver"
	messagesender "github.com/camopy/browser-chat/app/application/handler/messageSender"
	"github.com/camopy/browser-chat/app/application/service"
	rep "github.com/camopy/browser-chat/app/domain/repository"
	"github.com/camopy/browser-chat/app/infra/bot"
	"github.com/camopy/browser-chat/app/infra/broadcaster"
	"github.com/camopy/browser-chat/app/infra/mediator"
	"github.com/camopy/browser-chat/app/infra/repository"
	"github.com/camopy/browser-chat/app/infra/websocket"
)

func main() {
	db := repository.NewChatMessageMemoryRepository()
	mediator := mediator.New()
	broadcaster := broadcaster.New()
	bot := bot.New()
	registerEventHandlers(mediator, db, broadcaster, bot)
	startWebsocket(db, mediator, broadcaster)
}

func registerEventHandlers(mediator service.Mediator, repo rep.ChatMessageRepository, broadcaster service.Broadcaster, bot appbot.Bot) {
	senderHandler := messagesender.New(mediator, repo)
	mediator.Register(senderHandler)
	botMessageHandler := botmessagehandler.New(mediator, bot)
	mediator.Register(botMessageHandler)
	receiverHandler := messagereceiver.New(mediator, broadcaster)
	mediator.Register(receiverHandler)
}

func startWebsocket(db rep.ChatMessageRepository, mediator service.Mediator, broadcaster service.Broadcaster) {
	websocket := websocket.New(db, mediator, broadcaster)
	go websocket.HandleMessages(broadcaster)
	websocket.Start()
}
