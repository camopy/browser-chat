package bot

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
)

const name = "Gopher"

const quoteEndpoint = "https://stooq.com/q/l/?s=%s&f=sd2t2ohlcv&h&e=csv"

type Bot struct {
	name string
}

func New() *Bot {
	return &Bot{
		name: name,
	}
}

func (b *Bot) Name() string {
	return b.name
}

func (b *Bot) GetQuote(stockCode string) (string, error) {
	url := fmt.Sprintf(quoteEndpoint, stockCode)
	r, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	csv, err := readCsvFromBody(r.Body)
	if err != nil {
		return "", err
	}
	close := csv[1][6]
	return close, err
}

func readCsvFromBody(body io.ReadCloser) ([][]string, error) {
	reader := csv.NewReader(body)
	return reader.ReadAll()
}
