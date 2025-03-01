package main

import (
	"example.org/price-calculator/cmdmanager"
	"example.org/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		//manager := filemanager.New("prices.txt", fmt.Sprintf("resault_%.0f.json", taxRate*100))
		cmdManager := cmdmanager.New()
		taxJob := prices.NewTaxIncludedPriceJob(cmdManager, taxRate)
		taxJob.Process()
	}

}
