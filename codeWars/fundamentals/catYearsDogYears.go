/*
I have a cat and a dog.

I got them at the same time as kitten/puppy. That was humanYears years ago.
Return their respective ages now as [humanYears,catYears,dogYears]

NOTES:
humanYears >= 1
humanYears are whole numbers only
Cat Years
15 cat years for first year
+9 cat years for second year
+4 cat years for each year after that
Dog Years
15 dog years for first year
+9 dog years for second year
+5 dog years for each year after that
*/

package main

import "fmt"

func CalculateYears(years int) (result [3]int) {
	result[0] = years
	switch {
	case years == 1:
		result[1], result[2] = 15, 15
	case years == 2:
		result[1], result[2] = 24, 24
	case years > 2:
		result[1] = 24 + (years-2)*4
		result[2] = 24 + (years-2)*5
	}
	return
}

func main() {
	fmt.Println(CalculateYears(7))
}
