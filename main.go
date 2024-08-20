package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	results := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		taxIncludedPrices := make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxIncludedPrices[priceIndex] = price * (1 + taxRate)
		}
		results[taxRate] = taxIncludedPrices
	}

	for taxRate, prices := range results {
		fmt.Printf("Tax Rate: %.2f, Prices: %v\n", taxRate, prices)
	}

}
