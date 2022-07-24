package entity

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const stockCmd = "/stock="

var ErrMissingStockCode = errors.New("missing stock code")
var ErrUnknownBotCommand = errors.New("unknown bot command")

type BotMessage struct {
	BotName   string
	StockCode string
	Text      string
	Time      time.Time
}

func NewBotMessage(text string) (*BotMessage, error) {
	if !IsABotCommand(text) {
		return nil, ErrUnknownBotCommand
	}

	stockCode := extractStockCode(text)

	if stockCode == "" {
		return nil, ErrMissingStockCode
	}

	return &BotMessage{
		BotName:   "Gopher",
		StockCode: strings.ToLower(stockCode),
	}, nil
}

func extractStockCode(text string) string {
	return strings.Trim(strings.TrimPrefix(text, stockCmd), " ")
}

func IsABotCommand(text string) bool {
	return strings.HasPrefix(text, stockCmd)
}

func (b *BotMessage) Answer(close string) {
	b.Text = fmt.Sprintf("%s quote is $%s per share", strings.ToUpper(b.StockCode), close)
	b.Time = time.Now()
}
