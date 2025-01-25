package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("8.num")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	numStr := strings.ReplaceAll(strings.ReplaceAll(string(content), " ", ""), "\n", "")

	largest := 0
	for i := 0; i < len(numStr)-13; i++ {
		product := 1
		for j := 0; j < 13; j++ {
			product *= int(numStr[i+j] - '0')
		}
		if product > largest {
			largest = product
		}
	}
	fmt.Println("Largest product:", largest)
}
