package main

import (
	botmessagehandler "github.com/camopy/browser-chat/app/application/handler/botMessageHandler"
	messagereceiver "github.com/camopy/browser-chat/app/application/handler/messageReceiver"
	messagesender "github.com/camopy/browser-chat/app/application/handler/messageSender"
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

func registerEventHandlers(mediator *mediator.Mediator, repo *repository.ChatMessageRepository, broadcaster *broadcaster.Broadcaster, bot *bot.Bot) {
	senderHandler := messagesender.New(mediator, repo)
	mediator.Register(senderHandler)
	botMessageHandler := botmessagehandler.New(mediator, bot)
	mediator.Register(botMessageHandler)
	receiverHandler := messagereceiver.New(mediator, broadcaster)
	mediator.Register(receiverHandler)
}

func startWebsocket(db *repository.ChatMessageRepository, mediator *mediator.Mediator, broadcaster *broadcaster.Broadcaster) {
	websocket := websocket.New(db, mediator, broadcaster)
	go websocket.HandleMessages(broadcaster)
	websocket.Start()
}
