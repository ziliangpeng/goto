package main

import "fmt"

func main() {
	num := 600851475143
	var factors []int

	for i := 2; i <= num; i++ {
		for num%i == 0 {
			factors = append(factors, i)
			num /= i
		}
		if num == 1 {
			break
		}
	}
	fmt.Println(factors[len(factors)-1])
}
