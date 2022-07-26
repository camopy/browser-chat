package entity_test

import (
	"testing"

	"github.com/camopy/browser-chat/app/domain/entity"
	"github.com/stretchr/testify/assert"
)

type HashMock struct {
}

func (h HashMock) GeneratePassword(password string) (string, error) {
	if password == "password" {
		return "hashedpassword", nil
	}
	return "hashedpassword2", nil
}

func (h HashMock) CheckPassword(password, hash string) bool {
	return true
}

func TestNewUser(t *testing.T) {
	hash := &HashMock{}

	tests := []struct {
		name     string
		userName string
		password string
		want     *entity.User
	}{
		{
			name:     "user1",
			userName: "username",
			password: "password",
			want: &entity.User{
				UserName: "username",
				Password: "hashedpassword",
			},
		},
		{
			name:     "user2",
			userName: "username2",
			password: "password2",
			want: &entity.User{
				UserName: "username2",
				Password: "hashedpassword2",
			},
		},
		{
			name:     "empty user name",
			userName: "",
			password: "password",
			want:     nil,
		},
		{
			name:     "empty password",
			userName: "username",
			password: "",
			want:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := entity.NewUser(hash, tt.userName, tt.password)
			assert.Equal(t, tt.want, got)
			if tt.want != nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
