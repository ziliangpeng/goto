package main

func PrimesUnder(top int) []int {
	// TODO: explore sieve of eratosthenes
	primes := []int{}
	for i := 2; i < top; i++ {
		is_prime := true
		for _, p := range primes {
			if i%p == 0 {
				is_prime = false
				break
			}
		}
		if is_prime {
			primes = append(primes, i)
		}
	}
	return primes
}

func UniqueFactors(n int) []int {
	primes := PrimesUnder(n) // TODO: can use sqrt(n)
	factors := []int{}
	for _, p := range primes {
		if n%p == 0 {
			factors = append(factors, p)
		}
	}
	return factors
}
