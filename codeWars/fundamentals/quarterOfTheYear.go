/*
Given a month as an integer from 1 to 12, return to which quarter of the year it belongs as an integer number.

For example: month 2 (February), is part of the first quarter; month 6 (June), is part of the second quarter; and month 11 (November), is part of the fourth quarter.

Constraint:

1 <= month <= 12
*/

package main

import "fmt"

func QuarterOf(month int) (q int) {
	switch {
	case month < 4:
		q = 1
	case month >= 4 && month < 7:
		q = 2
	case month >= 7 && month < 10:
		q = 3
	case month >= 10 && month < 13:
		q = 4
	}
	return
}

func main() {
	fmt.Println(QuarterOf(2))
	fmt.Println(QuarterOf(5))
	fmt.Println(QuarterOf(9))
	fmt.Println(QuarterOf(12))
}
