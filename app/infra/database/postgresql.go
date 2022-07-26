package database

import (
	"database/sql"
	"fmt"

	"github.com/camopy/browser-chat/config"
	_ "github.com/lib/pq"
)

func NewPostgreSQL(conf *config.DbConf) (*sql.DB, error) {
	cs := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Port, conf.Host, conf.Username, conf.Password, conf.DbName)

	return sql.Open("postgres", cs)
}
