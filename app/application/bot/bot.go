package bot

type Bot interface {
	Name() string
	GetQuote(stockCode string) (string, error)
}
