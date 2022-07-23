package event

import "time"

const messageSubmitted = "MessageSubmitted"

type MessageSubmitted struct {
	name     string
	UserName string
	Message  string
	Time     time.Time
}

func NewMessageSubmitted(userName string, message string, time time.Time) *MessageSubmitted {
	return &MessageSubmitted{
		name:     messageSubmitted,
		UserName: userName,
		Message:  message,
		Time:     time,
	}
}

func (m *MessageSubmitted) Name() string {
	return m.name
}
