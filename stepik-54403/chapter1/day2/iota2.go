package main

import "fmt"

func main() {
	const (
		c0 = iota
		c1 = iota
		c2 = iota
	)
	fmt.Println(c0, c1, c2)

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		_ // 7
		Add
	)
	fmt.Println(Sunday)   // Output: 0
	fmt.Println(Saturday) // Output: 6
	fmt.Println(Add)      //// Output: 8

	const (
		u         = iota * 42 // u == 0
		v float64 = iota * 42 // v == 42.0 (iota= = 1, 1.0 * 42)
		w         = iota * 42 // w == 84 (iota == 2)
	)

	const x = iota // x == 0
	const y = iota // y == 0

	const (
		a     int     = 45
		b             // b == a
		c     float32 = 3.3
		d             // d == c
		First = iota  // iota == 4
		Second
		Third // Third == 6
	)
	fmt.Println(Third)

	const (
		_ = (iota + 2) * 2 // 4
		_                  // 6
		m                  // 8
		n                  // 10
		z                  // 12
	)
	fmt.Println(m, n, z)
}
