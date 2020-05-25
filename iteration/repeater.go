package main

import "fmt"

//Repeater repeat a character multiple times
func Repeater(char string, times int) string {
	var finalString string
	for index := 1; index <= times; index++ {
		finalString += char
	}
	return finalString
}

func main() {
	res := Repeater("a", 5)
	fmt.Println(res)
}
