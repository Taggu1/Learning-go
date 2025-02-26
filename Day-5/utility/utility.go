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

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
