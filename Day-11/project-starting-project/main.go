package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan error, len(taxRates), 4)

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index])

		// if err != nil {
		// }
	}

	for _, errChan := range doneChans {
		err := <-errChan
		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}
