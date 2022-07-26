package repository

import (
	passwordhash "github.com/camopy/browser-chat/app/application/service/passwordHash"
	"github.com/camopy/browser-chat/app/domain/entity"
	"github.com/camopy/browser-chat/app/domain/repository"
	"github.com/camopy/browser-chat/app/infra/database"
	"gorm.io/gorm"
)

type userRepository struct {
	hash passwordhash.PasswordHash
	*database.GORM
}

func NewUserRepository(hash passwordhash.PasswordHash, db *database.GORM) *userRepository {
	return &userRepository{hash, db}
}

type User struct {
	gorm.Model
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (ur *userRepository) CreateUser(userName, password string) (*entity.User, error) {
	u, err := entity.NewUser(ur.hash, userName, password)
	if err != nil {
		return nil, err
	}
	user := &User{
		UserName: u.UserName,
		Password: u.Password,
	}
	if err := ur.Create(user).Error; err != nil {
		return nil, err
	}
	u.ID = user.ID
	return u, nil
}

func (ur *userRepository) FindByUserName(userName string) (*entity.User, error) {
	user := &entity.User{}
	if err := ur.Where("user_name = ?", userName).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, repository.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}
