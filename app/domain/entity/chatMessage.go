package entity

import (
	"errors"
	"strings"
	"time"
)

var ErrMissingUserName = errors.New("missing userName")
var ErrMissingText = errors.New("missing text")
var ErrMissingTime = errors.New("missing time")

type ChatMessage struct {
	UserName string    `json:"userName"`
	Text     string    `json:"text"`
	Time     time.Time `json:"time"`
}

func NewChatMessage(userName string, text string, time time.Time) (*ChatMessage, error) {
	if strings.Trim(userName, " ") == "" {
		return nil, ErrMissingUserName
	}

	if strings.Trim(text, " ") == "" {
		return nil, ErrMissingText
	}

	if time.IsZero() {
		return nil, ErrMissingTime
	}

	return &ChatMessage{
		UserName: userName,
		Text:     text,
		Time:     time,
	}, nil
}
