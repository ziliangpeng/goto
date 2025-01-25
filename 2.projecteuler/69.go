package main

import "fmt"

// first, get factors for each n, using DP, and not start over for every i.
// then based on factors, maybe DFS number of non-coprime numbers. (maybe not)

func phi(n int) int {
	unique_factors := UniqueFactors(n)
	phi := 0
	for i := 1; i < n; i++ {
		is_coprime := true
		for _, factor := range unique_factors {
			if i%factor == 0 {
				is_coprime = false
				break
			}
		}
		if is_coprime {
			phi++
		}
	}
	return phi
}

func main() {
	const n = 1000000
	fmt.Println("n: ", n)

	max_totient := 0.
	best_n := 0

	for i := 2; i <= n; i++ {
		// the rationale is that n should be big; phi(n) should be small.
		// phi(n) small means not many coprimes = many numbers share factors = has many unique factors
		// = has factor of 2, 3, 5, 7, 11, 13, 17, ... preferrably.
		// and 2 * 3 * 5 * 7 * 11 * 13 * 17 = 510510
		if i%2 != 0 {
			continue
		}
		if i%3 != 0 {
			continue
		}
		if i%5 != 0 {
			continue
		}
		if i%7 != 0 {
			continue
		}
		if i%11 != 0 {
			continue
		}
		if i%13 != 0 {
			continue
		}
		phi_i := phi(i)
		// fmt.Println(i, phi_i, float64(i)/float64(phi_i))
		if float64(i)/float64(phi_i) > max_totient {
			max_totient = float64(i) / float64(phi_i)
			best_n = i
		}
	}

	fmt.Println("Max totient is: ", max_totient)
	fmt.Println(best_n)
}

// Answer is 510510
