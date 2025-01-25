package main

import "fmt"

func main() {
	var fib []int
	cap := 4000000
	fib = append(fib, 1, 2)
	for {
		next := fib[len(fib)-1] + fib[len(fib)-2]
		if next > cap {
			break
		}
		fib = append(fib, next)
	}
	fmt.Println(fib)

	sum := 0
	for _, v := range fib {
		if v%2 == 0 {
			sum += v
		}
	}
	fmt.Println("Answer is", sum)
}
