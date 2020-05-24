package main

import "fmt"

//Add two numbers and return the result
func Add(firstNum int, secondNum int) int {
	return firstNum + secondNum
}

func exampleAdd() {
	sum := Add(5, 7)
	fmt.Println(sum)
}

func main() {
	exampleAdd()
}
