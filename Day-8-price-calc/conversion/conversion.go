package conversion

import (
	"fmt"
	"strconv"
)

func ListOfStringsToFloats(list []string) []float64 {
	floatList := make([]float64, len(list))

	for index, priceString := range list {
		price, err := strconv.ParseFloat(priceString, 64)

		if err != nil {
			fmt.Println(err)
			continue
		}
		floatList[index] = price
	}

	return floatList
}
