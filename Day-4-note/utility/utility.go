package utility

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(text string) string {
	fmt.Print(text)
	reader := bufio.NewReader(os.Stdin)

	val, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.

	return val
}
