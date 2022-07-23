package entity_test

import (
	"testing"

	"github.com/camopy/browser-chat/app/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewChatMessage(t *testing.T) {
	tests := []struct {
		name     string
		userName string
		text     string
		want     *entity.ChatMessage
		err      error
	}{
		{
			name:     "success",
			userName: "userName",
			text:     "text",
			want: &entity.ChatMessage{
				UserName: "userName",
				Text:     "text",
			},
		},
		{
			name:     "success with space",
			userName: "userName2",
			text:     "text  text  text",
			want: &entity.ChatMessage{
				UserName: "userName2",
				Text:     "text  text  text",
			},
		},
		{
			name:     "failed - missing userName",
			userName: "  ",
			text:     "text",
			want:     &entity.ChatMessage{},
			err:      entity.ErrMissingUserName,
		},
		{
			name:     "failed - missing text",
			userName: "userName",
			text:     " ",
			want:     &entity.ChatMessage{},
			err:      entity.ErrMissingText,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := entity.NewChatMessage(tt.userName, tt.text)
			if tt.err != nil {
				assert.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want.Text, got.Text)
				assert.Equal(t, tt.want.UserName, got.UserName)
				assert.NotEmpty(t, got.Time)
			}
		})
	}
}
