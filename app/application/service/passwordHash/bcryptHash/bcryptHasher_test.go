package bcrypthash_test

import (
	"math/rand"
	"strings"
	"testing"

	bcrypthash "github.com/camopy/browser-chat/app/application/service/passwordHash/bcryptHash"
	"github.com/stretchr/testify/assert"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func TestGeneratePassword(t *testing.T) {
	h := bcrypthash.BcryptHash{}
	t.Run("no error", func(t *testing.T) {
		randomString := generateRandomString(16)
		got, err := h.GeneratePassword(randomString)

		assert.NoError(t, err)
		assert.NotEmpty(t, got)
	})

	t.Run("different outputs", func(t *testing.T) {
		randomString := generateRandomString(16)
		got1, err1 := h.GeneratePassword(randomString)
		got2, err2 := h.GeneratePassword(randomString)

		assert.NoError(t, err1)
		assert.NoError(t, err2)
		assert.NotEqual(t, got1, got2)
	})
}

func generateRandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func TestCheckPassword(t *testing.T) {
	h := bcrypthash.BcryptHash{}
	tests := []struct {
		name     string
		password string
		expected bool
	}{
		{
			name:     "correct password",
			password: "password",
			expected: true,
		},
		{
			name:     "incorrect password",
			password: "incorrect",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword, _ := h.GeneratePassword("password")
			actual := h.CheckPassword(tt.password, hashedPassword)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
