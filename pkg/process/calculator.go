package process

import (
	"math"
	"opg-analysis/internal/pos"
)

type calculator struct {
	// Profit percentage of the gap
	profitPercent float64
	// maximum amount we tolerate losing per trade
	maxLossPerTrade float64
}

func (c *calculator) Calculate(gapPercent, openingPrice float64) pos.Position {
	closingPrice := openingPrice / (1 + gapPercent)
	gapValue := closingPrice - openingPrice
	profitFromGap := c.profitPercent * gapValue

	stopLoss := openingPrice - profitFromGap
	takeProfit := openingPrice + profitFromGap

	shares := int(c.maxLossPerTrade / math.Abs(stopLoss-openingPrice))

	profit := math.Abs(openingPrice-takeProfit) * float64(shares)
	profit = math.Round(profit*100) / 100

	return pos.Position{
		EntryPrice:      math.Round(openingPrice*100) / 100,
		Shares:          shares,
		TakeProfitPrice: math.Round(takeProfit*100) / 100,
		StopLossPrice:   math.Round(stopLoss*100) / 100,
		Profit:          math.Round(profit*100) / 100,
	}
}

// NewCalculator returns a new instance of Calculator
//
// accountBalance is the amount of money in the trading account
// lossTolerance is the maximum percentage of the account balance we tolerate losing per trade
// profitPercent is the percentage of the gap we want to make as profit
func NewCalculator(accountBalance, lossTolerance, profitPercent float64) pos.Calculator {

	maxLossPerTrade := accountBalance * lossTolerance

	return &calculator{
		maxLossPerTrade: maxLossPerTrade,
		profitPercent:   profitPercent,
	}
}
