package main

import (
	"flag"
	"fmt"
	"opg-analysis/cmd"
	json2 "opg-analysis/pkg/json"
	"opg-analysis/pkg/my_csv"
	"opg-analysis/pkg/process"
	"opg-analysis/pkg/seeking_alpha"
	"os"
)

func main() {

	seekingAlphaURL := os.Getenv("SEEKING_ALPHA_URL")
	seekingAlphaAPIKey := os.Getenv("SEEKING_ALPHA_API_KEY")

	if seekingAlphaURL == "" {
		fmt.Println("Missing SEEKING_ALPHA_URL environment variable")
		os.Exit(1)
	}
	if seekingAlphaAPIKey == "" {
		fmt.Println("Missing SEEKING_ALPHA_API_KEY environment variable")
		os.Exit(1)
	}

	// Define command-line flags
	inputPath := flag.String("i", "", "path to input file (required)")
	accountBalance := flag.Float64("b", 0.0, "Account balance (required)")
	outputPath := flag.String("o", "./opg.json", "Path to output file.")
	lossTolerance := flag.Float64("l", 0.02, "Loss tolerance percentage")
	profitPercent := flag.Float64("p", 0.8, "Percentage of the gap to take as profit")
	minGap := flag.Float64("m", 0.1, "Minimum gap value to consider")

	// Parse command-line flags
	flag.Parse()

	// Check if required flags are provided
	if *inputPath == "" || *accountBalance == 0.0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var ldr = my_csv.NewLoader(*inputPath)
	var f = process.NewFilterer(*minGap)
	var c = process.NewCalculator(*accountBalance, *lossTolerance, *profitPercent)
	var fet = seeking_alpha.NewClient(seekingAlphaURL, seekingAlphaAPIKey)
	var del = json2.NewDeliverer(*outputPath)

	err := cmd.Run(ldr, f, c, fet, del)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//path := "./opg.my_csv"
	//stocks, err := Load(path)
	//
	//if err != nil {
	//	fmt.Println("Error loading stock data:", err)
	//	return
	//}
	//
	//stocks = slices.DeleteFunc(stocks, func(stock Stock) bool {
	//	//if math.Abs(stock.Gap) < .1 {
	//	//	return true
	//	//} else {
	//	//	return false
	//	//}
	//
	//	return math.Abs(stock.Gap) < .1
	//})
	//
	////var wg sync.WaitGroup
	//selectionChannel := make(chan Selection)
	//for _, stock := range stocks {
	//	//wg.Add(1)
	//	go func(stock Stock, selected chan<- Selection) {
	//		//defer wg.Done()
	//		position := Calculate(stock.Gap, stock.OpeningPrice)
	//
	//		articles, err := FetchNews(stock.Ticker)
	//		if err != nil {
	//			fmt.Printf("Error fetching news for %s: %v\n", stock.Ticker, err)
	//			selected <- Selection{}
	//			return
	//		} else {
	//			log.Printf("Found %d articles about %s", len(articles), stock.Ticker)
	//		}
	//
	//		selection := Selection{
	//			Ticker:   stock.Ticker,
	//			Position: position,
	//			Articles: articles,
	//		}
	//		selected <- selection
	//		//selections = append(selections, selection)
	//	}(stock, selectionChannel)
	//}
	//
	//var selections []Selection
	//for sel := range selectionChannel {
	//	selections = append(selections, sel)
	//	if len(selections) == len(stocks) {
	//		close(selectionChannel)
	//	}
	//}
	////wg.Wait()
	//
	//outputPath := "./opg.json"
	//err = Deliver(outputPath, selections)
	//if err != nil {
	//	fmt.Println("Error delivering selections to JSON:", err)
	//	return
	//}
	//fmt.Printf("Finished writing output  to %s\n", outputPath)
}
