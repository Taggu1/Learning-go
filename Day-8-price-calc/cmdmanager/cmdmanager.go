package cmdmanager

import "fmt"

type CMDManager struct{}

func (manager CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices confirm every price with ENTER.")
	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		_, err := fmt.Scan(&price)

		if err != nil {
			fmt.Println("Couldn't read price")
			continue
		}

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (manager CMDManager) WriteResault(data interface{}) error {
	_, err := fmt.Println(data)
	return err
}

func New() CMDManager {
	return CMDManager{}
}
