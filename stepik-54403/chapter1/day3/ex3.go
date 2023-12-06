/*
Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

Формат входных данных:
На вход дается натуральное число, не превосходящее 10000.

Формат выходных данных:
Выведите одно целое число - первую цифру заданного числа.
*/

package main

import "fmt"

func main() {
	var a uint16
	fmt.Scan((&a))

	switch {
	case a < 10:
		fmt.Println(a)
	case a < 100:
		fmt.Println(a / 10)
	case a < 1000:
		fmt.Println(a / 100)
	case a < 10000:
		fmt.Println(a / 1000)
	case a == 10000:
		fmt.Println(1)
	}

}
