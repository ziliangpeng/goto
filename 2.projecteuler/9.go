package main

import "fmt"

func main() {
	for a := 1; a < 1000; a++ {
		for b := a; b < 1000; b++ {
			c := 1000 - a - b
			if a*a+b*b == c*c {
				fmt.Println("a, b, c:", a, b, c)
				fmt.Println("Product:", a*b*c)
				return
			}
		}
	}

}
