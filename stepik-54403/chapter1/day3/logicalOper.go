package main

import "fmt"

func main() {
	a := true

	fmt.Println(!a)
	fmt.Println(26 > 21 && 6 > 1, 26 > 21 && 5 != 5)
	fmt.Println(26 < 21 || 6 < 1, 26 > 21 || 5 != 5)

}
