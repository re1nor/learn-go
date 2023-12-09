/*
Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
Программа в первой строке получает на вход число n - количество чисел в последовательности,
во второй строке -- n чисел, входящих в данную последовательность.
*/

package main

import "fmt"

func main() {
	var n, sum, num int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if num > 10 && num < 100 && num%8 == 0 {
			sum += num
		}
	}
	fmt.Println(sum)
}