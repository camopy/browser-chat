package entity

import (
	"fmt"

	passwordhash "github.com/camopy/browser-chat/app/application/service/passwordHash"
)

var ErrPasswordIsEmpty = fmt.Errorf("password is empty")
var ErrUserNameIsEmpty = fmt.Errorf("user name is empty")

type User struct {
	ID       uint
	UserName string
	Password string
}

func NewUser(hash passwordhash.PasswordHash, userName, password string) (*User, error) {
	if userName == "" {
		return nil, ErrUserNameIsEmpty
	}
	if password == "" {
		return nil, ErrPasswordIsEmpty
	}
	h, err := hash.GeneratePassword(password)
	if err != nil {
		return nil, err
	}
	return &User{
		UserName: userName,
		Password: h,
	}, nil
}
