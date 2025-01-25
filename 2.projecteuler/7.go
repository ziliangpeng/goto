package main

import "fmt"

func main() {
	primes := []int{}
	need_prime_cnt := 10001
	for i := 2; len(primes) < need_prime_cnt; i++ {
		is_prime := true
		for _, v := range primes {
			if i%v == 0 {
				is_prime = false
				break
			}
		}
		if is_prime {
			primes = append(primes, i)
		}
	}
	fmt.Println("Found ", len(primes), " primes")
	fmt.Println("The 10001st prime is ", primes[len(primes)-1])
}
