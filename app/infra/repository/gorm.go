package repository

import (
	"net/url"
	"time"

	"github.com/camopy/browser-chat/app/domain/entity"
	"github.com/camopy/browser-chat/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GORM struct {
	*gorm.DB
}

func NewGORM(conf *config.DbConf) (*GORM, error) {
	dsn := url.URL{
		User:     url.UserPassword(conf.Username, conf.Password),
		Scheme:   "postgres",
		Host:     conf.Host,
		Path:     conf.DbName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &GORM{db}, nil
}

type ChatMessage struct {
	gorm.Model
	UserName string
	Message  string
	SentAt   time.Time
}

func (r *GORM) CreateMessage(chatMessage *entity.ChatMessage) error {
	m := &ChatMessage{
		UserName: chatMessage.UserName,
		Message:  chatMessage.Text,
		SentAt:   chatMessage.Time,
	}
	return r.Create(m).Error
}
