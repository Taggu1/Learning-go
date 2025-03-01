package prices

import (
	"fmt"

	"example.org/price-calculator/conversion"
)

type IOManager interface {
	WriteResault(interface{}) error
	ReadLines() ([]string, error)
}

type TaxIncludedPriceJob struct {
	TaxRate           float64           `json:"tax_rate"`
	InputPrices       []float64         `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
	fm                IOManager         `json="-"`
}

func (job *TaxIncludedPriceJob) Process() {
	taxIncludedPrices := make(map[string]string, len(job.InputPrices))

	for _, price := range job.InputPrices {
		taxResault := price * (1 + job.TaxRate)
		taxIncludedPrices[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxResault)
	}

	job.TaxIncludedPrices = taxIncludedPrices
	job.fm.WriteResault(job)
}

func NewTaxIncludedPriceJob(fm IOManager, taxRate float64) *TaxIncludedPriceJob {
	inputPrices := ReadPricesFile(fm)

	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: inputPrices,
		fm:          fm,
	}

}

func ReadPricesFile(manager IOManager) []float64 {

	pricesStrings, err := manager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return []float64{}
	}

	prices := conversion.ListOfStringsToFloats(pricesStrings)

	return prices

}
