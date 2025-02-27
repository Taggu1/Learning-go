package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://www.google.com",
		"Aws":    "https://www.aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Aws"])
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)

}
