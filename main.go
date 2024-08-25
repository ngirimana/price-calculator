package main

import (
	"fmt"

	// cmdManager "example.com/price-calculator/cmdmanager"
	fileManager "example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := fileManager.New("prices.txt", fmt.Sprintf("result_%0.0f.json", taxRate*100))
		// cmdm := cmdManager.New()
		processJob := prices.NewTaxIncludePriceJob(fm, taxRate)
		go processJob.Process(doneChans[index], errorChans[index])
		// if err != nil {
		// 	fmt.Println("Could not process the job")
		// 	fmt.Print(err)
		// }
	}

	for index := range taxRates {
		select {

		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println("Error")
				fmt.Print(err)
			}

		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}

}
