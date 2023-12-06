package main

import "fmt"

func main() {
	var a, b, c = 9, 3, 3

	fmt.Println(b == c, a == c)
	fmt.Println(a > b, b > c)
	fmt.Println(b < a, b < c)
	fmt.Println(b <= a, a <= c)
	fmt.Println(a >= b, b >= c)
	fmt.Println(a != b, b != c)

}
