package main

import "fmt"

func main() {
	sum := 0

	for i := 1; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	j := 1
	for ; j < 10; j++ {
		fmt.Println(j * j)
	}

	for j < 10 {
		fmt.Println(j * j)
		j++
	}

	// while
	for j < 10 {
		fmt.Println(j * j)
		j++
	}

	/*
		infinity loop
		for {

		}
	*/

	// Scan data
	var n int
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		fmt.Println(n)
	}
}
