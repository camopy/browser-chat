package main

import (
	"log"

	appbot "github.com/camopy/browser-chat/app/application/bot"
	botmessagehandler "github.com/camopy/browser-chat/app/application/handler/botMessageHandler"
	messagereceiver "github.com/camopy/browser-chat/app/application/handler/messageReceiver"
	messagesender "github.com/camopy/browser-chat/app/application/handler/messageSender"
	"github.com/camopy/browser-chat/app/application/service"
	dr "github.com/camopy/browser-chat/app/domain/repository"
	"github.com/camopy/browser-chat/app/infra/bot"
	"github.com/camopy/browser-chat/app/infra/broadcaster"
	"github.com/camopy/browser-chat/app/infra/mediator"
	"github.com/camopy/browser-chat/app/infra/repository"
	"github.com/camopy/browser-chat/app/infra/websocket"
	"github.com/camopy/browser-chat/config"
)

func main() {
	config := config.AppConfig()
	db := initDb(&config.Db)
	mediator := mediator.New()
	broadcaster := broadcaster.New()
	bot := bot.New()
	registerEventHandlers(mediator, db, broadcaster, bot)
	startWebsocket(db, mediator, broadcaster, config)
}

func initDb(conf *config.DbConf) *repository.GORM {
	db, err := repository.NewGORM(conf)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	return db
}

func registerEventHandlers(mediator service.Mediator, repo dr.ChatMessageRepository, broadcaster service.Broadcaster, bot appbot.Bot) {
	senderHandler := messagesender.New(mediator, repo)
	mediator.Register(senderHandler)
	botMessageHandler := botmessagehandler.New(mediator, bot)
	mediator.Register(botMessageHandler)
	receiverHandler := messagereceiver.New(mediator, broadcaster)
	mediator.Register(receiverHandler)
}

func startWebsocket(db dr.ChatMessageRepository, mediator service.Mediator, broadcaster service.Broadcaster, conf *config.Conf) {
	websocket := websocket.New(db, mediator, broadcaster, conf)
	go websocket.HandleMessages(broadcaster)
	websocket.Start()
}
