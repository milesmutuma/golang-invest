// Package trade provides a trade selection type and a deliverer
// interface for delivering trade selections.

package trade

import (
	"opg-analysis/internal/news"
	"opg-analysis/internal/pos"
)

// Selection represents a trade selection, including the ticker symbol,
// position information, and a list of related news articles.
type Selection struct {
	Ticker   string
	Position pos.Position
	Articles []news.Article
}

// Deliverer is an interface for delivering trade selections.
type Deliverer interface {
	Deliver(selections []Selection) error
}
