package main

import (
	"fmt"

	// cmdManager "example.com/price-calculator/cmdmanager"
	fileManager "example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := fileManager.New("prices.txt", fmt.Sprintf("result_%0.0f.json", taxRate*100))
		// cmdm := cmdManager.New()
		processJob := prices.NewTaxIncludePriceJob(fm, taxRate)
		err := processJob.Process()
		if err != nil {
			fmt.Println("Could not process the job")
			fmt.Print(err)
		}
	}

}
