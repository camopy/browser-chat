package usersignup_test

import (
	"testing"

	bcrypthash "github.com/camopy/browser-chat/app/application/service/passwordHash/bcryptHash"
	usersignup "github.com/camopy/browser-chat/app/application/usecase/userSignup"
	"github.com/camopy/browser-chat/app/infra/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserSignup(t *testing.T) {
	hash := &bcrypthash.BcryptHash{}
	repo := repository.NewUserRepositoryMemory(hash)
	userSignup := usersignup.NewUserSignup(repo)
	t.Run("should return error if user name is empty", func(t *testing.T) {
		input := usersignup.Input{
			UserName: "",
			Password: "password",
		}
		_, err := userSignup.Execute(input)
		assert.Error(t, err)
	})

	t.Run("should return error if password is empty", func(t *testing.T) {
		input := usersignup.Input{
			UserName: "username",
			Password: "",
		}
		_, err := userSignup.Execute(input)
		assert.Error(t, err)
	})

	t.Run("should return error if user name already exists", func(t *testing.T) {
		input := usersignup.Input{
			UserName: "username",
			Password: "password",
		}
		_, err := userSignup.Execute(input)
		require.NoError(t, err)
		_, err = userSignup.Execute(input)
		assert.Error(t, err)
	})

	t.Run("should return user id and user name", func(t *testing.T) {
		input := usersignup.Input{
			UserName: "username2",
			Password: "password",
		}
		output, err := userSignup.Execute(input)
		require.NoError(t, err)
		assert.Equal(t, "username2", output.UserName)
	})
}
