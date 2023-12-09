/*
Write a function named setAlarm/set_alarm/set-alarm/setalarm (depending on language) which receives two parameters. The first parameter, employed,
is true whenever you are employed and the second parameter, vacation is true whenever you are on vacation.

The function should return true if you are employed and not on vacation (because these are the circumstances under which you need to set an alarm).
It should return false otherwise. Examples:

employed | vacation
true     | true     => false
true     | false    => true
false    | true     => false
false    | false    => false
*/

package main

import "fmt"

func SetAlarm(employed, vacation bool) bool {
	switch {
	case employed != vacation && vacation != true:
		return true
	default:
		return false
	}
}

func main() {
	fmt.Println(SetAlarm(true, true))
	fmt.Println(SetAlarm(true, false))
	fmt.Println(SetAlarm(false, true))
	fmt.Println(SetAlarm(false, false))
}
