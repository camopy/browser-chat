package repository

import "github.com/camopy/browser-chat/app/domain/entity"

type ChatMessageRepository struct {
	messages []*entity.ChatMessage
}

func NewChatMessageMemoryRepository() *ChatMessageRepository {
	return &ChatMessageRepository{
		messages: []*entity.ChatMessage{},
	}
}

func (r *ChatMessageRepository) CreateMessage(chatMessage *entity.ChatMessage) error {
	r.messages = append(r.messages, chatMessage)
	return nil
}
