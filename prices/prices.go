package prices

import (
	"bufio"
	"fmt"
	"os"
)

type TaxIncludePriceJob struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludePrices map[string]float64
}

func (job TaxIncludePriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("cannot open file")
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading the file failed.")
		fmt.Println(err)
		file.Close()
		return
	}
	

}
func (job TaxIncludePriceJob) Process() {
	results := make(map[string]float64)

	for _, price := range job.InputPrices {
		results[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	fmt.Println(results)
}

func NewTaxIncludePriceJob(taxRate float64) *TaxIncludePriceJob {
	return &TaxIncludePriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
