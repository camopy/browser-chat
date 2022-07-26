package repository

import (
	"fmt"

	"github.com/camopy/browser-chat/app/domain/entity"
)

var ErrUserNotFound = fmt.Errorf("user not found")
var ErrUserNameAlreadyExists = fmt.Errorf("user name already exists")

type UserRepository interface {
	FindByUserName(userName string) (*entity.User, error)
	CreateUser(userName, password string) (*entity.User, error)
}
