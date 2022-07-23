package event

import "time"

const messageSent = "MessageSent"

type MessageSent struct {
	name     string
	UserName string
	Message  string
	Time     time.Time
}

func NewMessageSent(userName string, message string, time time.Time) *MessageSent {
	return &MessageSent{
		name:     messageSent,
		UserName: userName,
		Message:  message,
		Time:     time,
	}
}

func (m *MessageSent) Name() string {
	return m.name
}
