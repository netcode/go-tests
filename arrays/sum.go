package main

import "fmt"

//Sum of 5 numbers
func Sum(numbers [5]int) int {
	var total int
	for _, i := range numbers {
		total += i
	}
	return total
}

func main() {
	fmt.Println("Hello Sum")
}
