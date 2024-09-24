// Package pos provides functionality for calculating the trading position.
package pos

// Position represents a trading position with its entry price, number of shares, take profit price,
// stop loss price, and profit.
type Position struct {
	// The price at which to buy or sell
	EntryPrice float64

	// How many shares to buy or sell
	Shares int

	// The price at which to exit and take my profit
	TakeProfitPrice float64

	// the price at which to stop my loss if the stock doesn’t go our way
	StopLossPrice float64

	// Expected final Profit
	Profit float64
}

// Calculator is an interface that defines a method for calculating a trading position based on the
// gap percentage and opening price.

type Calculator interface {
	Calculate(gapPercent, openingPrice float64) Position
}
