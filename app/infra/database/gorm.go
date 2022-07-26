package database

import (
	"net/url"

	"github.com/camopy/browser-chat/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GORM struct {
	*gorm.DB
}

func NewGORM(conf config.DbConf) (*GORM, error) {
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
