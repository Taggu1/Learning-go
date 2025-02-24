package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string) float64 {
	floatText, err := os.ReadFile(fileName)

	if err != nil {
		WriteFloatToFile(0, fileName)
		return 0
	}

	float, err := strconv.ParseFloat(string(floatText), 64)

	if err != nil {
		return 0
	}

	return float
}
