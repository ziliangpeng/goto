package main

import "fmt"

func main() {
	var num int
	for num = 1; true; num++ {
		ok := true
		for i := 1; i <= 20; i++ {
			if num%i != 0 {
				ok = false
				break
			}
		}
		if ok {
			break
		}
	}
	fmt.Println("Answer is", num)
}
