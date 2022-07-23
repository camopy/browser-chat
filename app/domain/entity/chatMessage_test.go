package entity_test

import (
	"testing"
	"time"

	"github.com/camopy/browser-chat/app/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewChatMessage(t *testing.T) {
	now := time.Now()
	tomorrow := now.Add(24 * time.Hour)
	tests := []struct {
		name     string
		userName string
		text     string
		time     time.Time
		want     *entity.ChatMessage
		err      error
	}{
		{
			name:     "success",
			userName: "userName",
			text:     "text",
			time:     now,
			want: &entity.ChatMessage{
				UserName: "userName",
				Text:     "text",
				Time:     now,
			},
		},
		{
			name:     "success with space",
			userName: "userName2",
			text:     "text  text  text",
			time:     tomorrow,
			want: &entity.ChatMessage{
				UserName: "userName2",
				Text:     "text  text  text",
				Time:     tomorrow,
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
		{
			name:     "failed - missing time",
			userName: "userName",
			text:     "message ",
			want:     &entity.ChatMessage{},
			err:      entity.ErrMissingTime,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := entity.NewChatMessage(tt.userName, tt.text, tt.time)
			if tt.err != nil {
				assert.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
