package entity

import (
	"errors"
	"strings"
	"time"
)

var ErrMissingUserName = errors.New("missing userName")
var ErrMissingText = errors.New("missing text")

type ChatMessage struct {
	UserName string    `json:"userName"`
	Text     string    `json:"text"`
	Time     time.Time `json:"time"`
}

func NewChatMessage(userName string, text string) (*ChatMessage, error) {
	if strings.Trim(userName, " ") == "" {
		return nil, ErrMissingUserName
	}

	if strings.Trim(text, " ") == "" {
		return nil, ErrMissingText
	}

	return &ChatMessage{
		UserName: userName,
		Text:     text,
		Time:     time.Now(),
	}, nil
}
