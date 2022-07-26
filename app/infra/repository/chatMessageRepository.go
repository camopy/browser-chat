package repository

import (
	"time"

	"github.com/camopy/browser-chat/app/domain/entity"
	"github.com/camopy/browser-chat/app/infra/database"
	"gorm.io/gorm"
)

type chatMessageRepository struct {
	*database.GORM
}

func NewChatMessageRepository(db *database.GORM) *chatMessageRepository {
	return &chatMessageRepository{db}
}

type ChatMessage struct {
	gorm.Model
	UserName string
	Message  string
	SentAt   time.Time
}

func (cr *chatMessageRepository) CreateMessage(chatMessage *entity.ChatMessage) error {
	m := &ChatMessage{
		UserName: chatMessage.UserName,
		Message:  chatMessage.Text,
		SentAt:   chatMessage.Time,
	}
	return cr.Create(m).Error
}
