package main

import "fmt"

func main() {
	var hello string
	hello = "Hello, Go!"
	var goodbye = "Goodbye"
	var a, b string
	var (
		name string = "Bulat"
		age  int    = 21
	)
	var z float64 = 1.2626
	var symbol int32 = 'b'
	c := 3

	fmt.Println(hello + goodbye)
	fmt.Println(a, b)
	fmt.Println("My names is", name)
	fmt.Println("I'm", age, "years old")
	fmt.Println(z)
	fmt.Println(symbol)
	fmt.Println(c)
}
