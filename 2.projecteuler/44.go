package main

import (
	"fmt"
)

func main() {
	const num = 10000
	pents := []int{}
	for i := 1; i < num; i++ {
		pents = append(pents, i*(3*i-1)/2)
	}
	pentSet := make(map[int]bool)
	for _, p := range pents {
		pentSet[p] = true
	}

	min_diff := 1000000000
	ans_a, ans_b := 0, 0
	for i, a := range pents {
		for _, b := range pents[i+1:] {
			sum := a + b
			diff := b - a
			if pentSet[sum] && pentSet[diff] {
				fmt.Println("a, b:", a, b)
				fmt.Println("Difference:", diff)
				fmt.Println("Sum:", sum)
				if diff < min_diff {
					min_diff = diff
					ans_a, ans_b = a, b
				}
			}
		}
	}
	fmt.Println("Min difference:", min_diff)
	fmt.Println("a, b:", ans_a, ans_b)
	ans := func() int {
		if ans_a > ans_b {
			return ans_a - ans_b
		}
		return ans_b - ans_a
	}()
	fmt.Println("Answer:", ans)
}
