/*
Introduction:
The first century spans from the year 1 up to and including the year 100, the second century - from the year 101 up to and including the year 200, etc.

Task:
Given a year, return the century it is in.

Examples:
1705 --> 18
1900 --> 19
1601 --> 17
2000 --> 20
*/

package main

import "fmt"

func century(year int) int {
	if (year % 100) != 0 {
		cent := year/100 + 1
		return cent
	} else {
		return year / 100
	}
}

func main() {
	out := century(2001)
	fmt.Println(out)
}
