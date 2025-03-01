package main

import (
	"example.org/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		taxJob := prices.NewTaxIncludedPriceJob(taxRate)
		taxJob.Process()
	}

}
