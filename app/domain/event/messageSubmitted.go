package event

import "time"

const name = "MessageSubmitted"

type MessageSubmitted struct {
	name     string
	UserName string
	Message  string
	Time     time.Time
}

func NewMessageSubmitted(userName string, message string, time time.Time) *MessageSubmitted {
	return &MessageSubmitted{
		name:     name,
		UserName: userName,
		Message:  message,
		Time:     time,
	}
}

func (m *MessageSubmitted) Name() string {
	return m.name
}
