package repository

import "github.com/camopy/browser-chat/app/domain/entity"

type ChatMessageRepository interface {
	CreateMessage(chatMessage *entity.ChatMessage) error
}
