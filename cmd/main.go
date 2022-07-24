package main

import (
	messagereceiver "github.com/camopy/browser-chat/app/application/handler/messageReceiver"
	messagesender "github.com/camopy/browser-chat/app/application/handler/messageSender"
	"github.com/camopy/browser-chat/app/infra/broadcaster"
	"github.com/camopy/browser-chat/app/infra/mediator"
	"github.com/camopy/browser-chat/app/infra/repository"
	"github.com/camopy/browser-chat/app/infra/websocket"
)

func main() {
	db := repository.NewChatMessageMemoryRepository()
	mediator := mediator.New()
	broadcaster := broadcaster.New()
	registerEventHandlers(mediator, db, broadcaster)
	startWebsocket(db, mediator, broadcaster)
}

func registerEventHandlers(mediator *mediator.Mediator, repo *repository.ChatMessageRepository, broadcaster *broadcaster.Broadcaster) {
	senderHandler := messagesender.New(mediator, repo)
	mediator.Register(senderHandler)
	receiverHandler := messagereceiver.New(mediator, broadcaster)
	mediator.Register(receiverHandler)
}

func startWebsocket(db *repository.ChatMessageRepository, mediator *mediator.Mediator, broadcaster *broadcaster.Broadcaster) {
	websocket := websocket.New(db, mediator, broadcaster)
	go websocket.HandleMessages(broadcaster)
	websocket.Start()
}
