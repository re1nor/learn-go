/*
Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.
Входные данные: На вход программе подается целое число d (0 < d < 360).
Выведите на экран фразу: It is ... hours ... minutes.
Вместо многоточий программа должна выводить значения h и m, отделяя их от слов ровно одним пробелом.
*/

package main

import "fmt"

func main() {
	var a uint
	fmt.Scan(&a)
	fmt.Println("It is", a/30, "hours", a%30*2, "minutes.")
}