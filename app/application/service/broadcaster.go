package service

import "github.com/camopy/browser-chat/app/domain/entity"

type Broadcaster interface {
	Broadcast(*entity.ChatMessage)
	Receive() *entity.ChatMessage
}
