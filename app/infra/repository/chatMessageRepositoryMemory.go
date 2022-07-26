package repository

import "github.com/camopy/browser-chat/app/domain/entity"

type ChatMessageRepositoryMemory struct {
	messages []*entity.ChatMessage
}

func NewChatMessageMemoryRepository() *ChatMessageRepositoryMemory {
	return &ChatMessageRepositoryMemory{
		messages: []*entity.ChatMessage{},
	}
}

func (r *ChatMessageRepositoryMemory) CreateMessage(chatMessage *entity.ChatMessage) error {
	r.messages = append(r.messages, chatMessage)
	return nil
}
