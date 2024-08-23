package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	ioManager "example.com/price-calculator/iomanager"
)

type TaxIncludePriceJob struct {
	IOManager        ioManager.IOManager `json:"-"`
	TaxRate          float64             `json:"tax_rate"`
	InputPrices      []float64           `json:"input_prices"`
	TaxIncludePrices map[string]string   `json:"tax_include_prices"`
}

func (job *TaxIncludePriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return err
	}

	prices, err := conversion.StringToFloat(lines)

	if err != nil {

		fmt.Println(err)
		return err
	}
	job.InputPrices = prices
	return nil

}

func (job *TaxIncludePriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}
	results := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		results[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludePrices = results

	return job.IOManager.WriteJSON(job)

}

func NewTaxIncludePriceJob(iom ioManager.IOManager, taxRate float64) *TaxIncludePriceJob {
	return &TaxIncludePriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
