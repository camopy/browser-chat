package entity_test

import (
	"testing"

	"github.com/camopy/browser-chat/app/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBotMessage(t *testing.T) {
	tests := []struct {
		name string
		text string
		want *entity.BotMessage
		err  error
	}{
		{
			name: "success",
			text: "/stock=CODE",
			want: &entity.BotMessage{
				BotName:   "Gopher",
				StockCode: "code",
			},
		},
		{
			name: "success with space at end",
			text: "/stock=Code2 ",
			want: &entity.BotMessage{
				BotName:   "Gopher",
				StockCode: "code2",
			},
		},
		{
			name: "failed - missing text",
			text: " ",
			want: &entity.BotMessage{},
			err:  entity.ErrUnknownBotCommand,
		},
		{
			name: "failed - missing stock code",
			text: "/stock= ",
			want: &entity.BotMessage{},
			err:  entity.ErrMissingStockCode,
		},
		{
			name: "failed - wrong command",
			text: "/stok=code",
			want: &entity.BotMessage{},
			err:  entity.ErrUnknownBotCommand,
		},
		{
			name: "failed - wrong command",
			text: "stock=code",
			want: &entity.BotMessage{},
			err:  entity.ErrUnknownBotCommand,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := entity.NewBotMessage(tt.text)
			if tt.err != nil {
				assert.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestAnswer(t *testing.T) {
	msg1, _ := entity.NewBotMessage("/stock=code")
	msg2, _ := entity.NewBotMessage("/stock=code2")

	tests := []struct {
		name          string
		msg           *entity.BotMessage
		wantedBotName string
		wantedClose   string
		wantedText    string
	}{
		{
			name:          "code",
			msg:           msg1,
			wantedBotName: "Gopher",
			wantedClose:   "123.45",
			wantedText:    "CODE quote is $123.45 per share",
		},
		{
			name:          "code2",
			msg:           msg2,
			wantedBotName: "Gopher",
			wantedClose:   "133.45",
			wantedText:    "CODE2 quote is $133.45 per share",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.msg.Answer(tt.wantedClose)
			assert.Equal(t, tt.wantedBotName, tt.msg.BotName)
			assert.Equal(t, tt.wantedText, tt.msg.Text)
			assert.NotEmpty(t, tt.msg.Time)
		})
	}
}
