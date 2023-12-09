/*
Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности,
которые равны ее наибольшему элементу.

Формат входных данных:
Вводится непустая последовательность натуральных чисел, оканчивающаяся числом 0 (само число 0 в последовательность не входит,
а служит как признак ее окончания).

Формат выходных данных:
Выведите ответ на задачу.
*/

package main

import "fmt"

func main() {
	var n, k, max int
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		if n > max {
			max = n
			k = 1
		} else if n == max {
			k++
		}

	}
	fmt.Println(k)
}