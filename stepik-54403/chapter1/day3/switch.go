package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
		fallthrough // if i==3 Print 3 4
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}

	switch {
	case 1 <= i && i <= 9:
		fmt.Println("от 1 до 9")
	case 100 <= i && i <= 250:
		fmt.Println("от 100 до 250")
	case 1000 <= i && i <= 6000:
		fmt.Println("от 1000 до 6000")
	}
}
