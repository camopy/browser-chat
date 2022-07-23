package submitmessage_test

import (
	"testing"
	"time"

	messagereceiver "github.com/camopy/browser-chat/app/application/handler/messageReceiver"
	messagesender "github.com/camopy/browser-chat/app/application/handler/messageSender"
	submitmessage "github.com/camopy/browser-chat/app/application/usecase/submitMessage"
	"github.com/camopy/browser-chat/app/infra/broadcaster"
	"github.com/camopy/browser-chat/app/infra/mediator"
	"github.com/camopy/browser-chat/app/infra/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubmitMessage(t *testing.T) {
	mediator := mediator.New()
	repo := repository.NewChatMessageMemoryRepository()
	senderHandler := messagesender.New(mediator, repo)
	mediator.Register(senderHandler)

	broadcaster := broadcaster.New()
	receiverHandler := messagereceiver.New(mediator, broadcaster)
	mediator.Register(receiverHandler)

	submit := submitmessage.New(mediator)
	now := time.Now()
	input := submitmessage.Input{
		UserName: "user name",
		Message:  "message",
		Time:     now,
	}
	submit.Execute(input)

	t.Run("should have one message stored on database", func(t *testing.T) {
		got, err := repo.FindAll()
		require.NoError(t, err)
		assert.Equal(t, 1, len(got))
		assert.Equal(t, input.UserName, got[0].UserName)
		assert.Equal(t, input.Message, got[0].Text)
		assert.Equal(t, input.Time, got[0].Time)
	})

	t.Run("should have one message broadcasted", func(t *testing.T) {
		receivedMsg := broadcaster.Receive()
		assert.Equal(t, "user name", receivedMsg.UserName)
		assert.Equal(t, "message", receivedMsg.Text)
		assert.Equal(t, now, receivedMsg.Time)
	})

}