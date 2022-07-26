package repository

import (
	passwordhash "github.com/camopy/browser-chat/app/application/service/passwordHash"
	"github.com/camopy/browser-chat/app/domain/entity"
	"github.com/camopy/browser-chat/app/domain/repository"
)

type UserRepositoryMemory struct {
	users map[string]*entity.User
	hash  passwordhash.PasswordHash
}

func NewUserRepositoryMemory(hash passwordhash.PasswordHash) *UserRepositoryMemory {
	return &UserRepositoryMemory{
		users: map[string]*entity.User{},
		hash:  hash,
	}
}

func (r *UserRepositoryMemory) FindByUserName(userName string) (*entity.User, error) {
	if user, ok := r.users[userName]; ok {
		return user, nil
	}
	return nil, repository.ErrUserNotFound
}

func (r *UserRepositoryMemory) CreateUser(userName, password string) (*entity.User, error) {
	if user, ok := r.users[userName]; ok {
		return user, nil
	}
	user, err := entity.NewUser(r.hash, userName, password)
	if err != nil {
		return nil, err
	}
	r.users[userName] = user
	return user, nil
}
