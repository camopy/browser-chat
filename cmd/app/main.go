package main

import (
	"log"

	botmessagehandler "github.com/camopy/browser-chat/app/application/handler/botMessageHandler"
	messagereceiver "github.com/camopy/browser-chat/app/application/handler/messageReceiver"
	messagesender "github.com/camopy/browser-chat/app/application/handler/messageSender"
	"github.com/camopy/browser-chat/app/application/service"
	bcrypthash "github.com/camopy/browser-chat/app/application/service/passwordHash/bcryptHash"
	dr "github.com/camopy/browser-chat/app/domain/repository"
	"github.com/camopy/browser-chat/app/infra/bot"
	"github.com/camopy/browser-chat/app/infra/broadcaster"
	"github.com/camopy/browser-chat/app/infra/database"
	"github.com/camopy/browser-chat/app/infra/mediator"
	"github.com/camopy/browser-chat/app/infra/repository"
	"github.com/camopy/browser-chat/app/infra/websocket"
	"github.com/camopy/browser-chat/app/util/validator"
	"github.com/camopy/browser-chat/config"
	goval "github.com/go-playground/validator"
)

func main() {
	config := config.AppConfig()
	db := initDb(config.Db)
	startChat(db, config.Server)
}

func initDb(conf config.DbConf) *database.GORM {
	db, err := database.NewGORM(conf)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	return db
}

func startChat(db *database.GORM, config config.ServerConf) {
	mediator := mediator.New()
	broadcaster := broadcaster.New()
	chatRepository := repository.NewChatMessageRepository(db)
	registerEventHandlers(mediator, chatRepository, broadcaster)
	hash := &bcrypthash.BcryptHash{}
	userRepository := repository.NewUserRepository(hash, db)
	validator := validator.New()
	startWebsocket(chatRepository, userRepository, mediator, broadcaster, config, validator)
}

func registerEventHandlers(mediator service.Mediator, chatRepo dr.ChatMessageRepository, broadcaster service.Broadcaster) {
	bot := bot.New()
	senderHandler := messagesender.New(mediator, chatRepo)
	mediator.Register(senderHandler)
	botMessageHandler := botmessagehandler.New(mediator, bot)
	mediator.Register(botMessageHandler)
	receiverHandler := messagereceiver.New(mediator, broadcaster)
	mediator.Register(receiverHandler)
}

func startWebsocket(chatRepo dr.ChatMessageRepository, userRepo dr.UserRepository, mediator service.Mediator, broadcaster service.Broadcaster, conf config.ServerConf, validator *goval.Validate) {
	websocket := websocket.New(chatRepo, userRepo, mediator, broadcaster, conf, validator)
	go websocket.HandleMessages(broadcaster)
	websocket.Start()
}
