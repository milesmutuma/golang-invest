package raw

type Stock struct {
	Ticker       string
	Gap          float64
	OpeningPrice float64
}

// Loader is an interface for loading data.
type Loader interface {
	Load() ([]Stock, error)
}

// Filterer is an interface for filtering raw stock data
type Filterer interface {
	Filter(stocks []Stock) []Stock
}
