package submitmessage_test

import (
	"testing"
	"time"

	messagesender "github.com/camopy/browser-chat/app/application/handler/messageSender"
	submitmessage "github.com/camopy/browser-chat/app/application/usecase/submitMessage"
	"github.com/camopy/browser-chat/app/infra/mediator"
	"github.com/camopy/browser-chat/app/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestSubmitMessage(t *testing.T) {
	mediator := mediator.New()
	repo := repository.NewChatMessageMemoryRepository()
	senderHandler := messagesender.New(mediator, repo)
	mediator.Register(senderHandler)

	submit := submitmessage.New(mediator)
	input := submitmessage.Input{
		UserName: "user name",
		Message:  "message",
		Time:     time.Now(),
	}
	submit.Execute(input)

	messages, _ := repo.FindAll()
	assert.Equal(t, 1, len(messages))
	assert.Equal(t, "user name", messages[0].UserName)
	assert.Equal(t, "message", messages[0].Text)
}
