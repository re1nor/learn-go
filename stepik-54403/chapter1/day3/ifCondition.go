package main

import "fmt"

func main() {
	a, b := 3, 6
	if a < b {
		fmt.Println("a<b")
	}
	c := 10
	if v := a + b; v < c {
		fmt.Println("v<c")
	}
}
