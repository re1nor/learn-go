package main

import "fmt"

func main() {
	sum := 0

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // next iter
		}
		sum += i
	}
	fmt.Println("Sum: ", sum)

	sum2 := 0
	for i := 1; i <= 9; i++ {
		if i > 4 {
			break // exit from loop
		}
		sum2 += i
	}
	fmt.Println("Sum2: ", sum2)
}
