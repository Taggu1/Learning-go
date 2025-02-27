package lists

import "fmt"

// import "fmt"

// func main() {
// 	prices := []float64{1.2, 2.4, 4, 2.4}

// 	featuredPrice := prices[1:]
// 	highlightedPirce := featuredPrice[:1]
// 	featuredPrice[0] = 22
// 	fmt.Println(prices)

// 	fmt.Println(len(featuredPrice), cap(featuredPrice))
// 	fmt.Println(highlightedPirce)
// }

func main() {
	prices := []float64{10.2, 12.3}
	fmt.Println(len(prices), cap(prices))
	prices = append(prices, 23.4)
	fmt.Println(len(prices), cap(prices))

	discountedPrices := []float64{1.2, 4.2}
	prices = append(prices, discountedPrices...)

	fmt.Println(prices)
}
