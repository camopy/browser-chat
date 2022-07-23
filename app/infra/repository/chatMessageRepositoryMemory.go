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

func (r *ChatMessageRepository) Save(chatMessage *entity.ChatMessage) error {
	r.messages = append(r.messages, chatMessage)
	return nil
}

func (r *ChatMessageRepository) FindAll() ([]*entity.ChatMessage, error) {
	return r.messages, nil
}
