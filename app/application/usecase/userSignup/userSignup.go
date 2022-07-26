package usersignup

import (
	"github.com/camopy/browser-chat/app/domain/repository"
)

type Input struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type Output struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
}

type userSignup struct {
	userRepository repository.UserRepository
}

func NewUserSignup(userRepository repository.UserRepository) *userSignup {
	return &userSignup{
		userRepository: userRepository,
	}
}

func (u *userSignup) Execute(input Input) (*Output, error) {
	_, err := u.userRepository.FindByUserName(input.UserName)
	if err == nil {
		return nil, repository.ErrUserNameAlreadyExists
	}
	user, err := u.userRepository.CreateUser(input.UserName, input.Password)
	if err != nil {
		return nil, err
	}

	return &Output{
		ID:       user.ID,
		UserName: user.UserName,
	}, nil
}
