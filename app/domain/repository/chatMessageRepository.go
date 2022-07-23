package repository

import "github.com/camopy/browser-chat/app/domain/entity"

type ChatMessageRepository interface {
	Save(chatMessage *entity.ChatMessage) error
	FindAll() ([]*entity.ChatMessage, error)
}
