package main

import "fmt"

func main() {
	a := 101
	b := 10

	// Addition
	fmt.Println("Result of a+b=", a+b)

	// Subtraction
	fmt.Println("Result of a-b=", a-b)

	// Multiplication
	fmt.Println("Result of a*b=", a*b)

	// Division
	fmt.Println("Int Result of a/b=", a/b)
	fmt.Println("For float", float32(a)/float32(b))

	// Modulus
	fmt.Println("Result of a%b=", a%b)

	// Increment
	a++
	fmt.Println("Result of a++ ", a)

	// Decrement
	b--
	fmt.Println("Result of a++ ", b)

}
