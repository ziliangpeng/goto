package main

import (
	"fmt"
	"sort"
	"strconv"
)

func isPalindrome(num int) bool {
	numStr := strconv.Itoa(num)
	return numStr == reverse(numStr)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	products := []int{}

	for i := 999; i > 99; i-- {
		for j := i; j > 99; j-- {
			products = append(products, i*j)
		}
	}
	// sort the products in descending order
	sort.Slice(products, func(a, b int) bool {
		return products[a] > products[b]
	})

	for _, v := range products {
		if isPalindrome(v) {
			fmt.Println(v)
			break
		}
	}
}
