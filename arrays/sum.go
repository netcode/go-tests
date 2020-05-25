package main

import "fmt"

//Sum of any numbers
func Sum(numbers []int) int {
	var total int
	for _, i := range numbers {
		total += i
	}
	return total
}

//SumAll arrays and return array with the result
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

//SumAllTails arrays and return array with the result
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}
		sums = append(sums, Sum(numbers[1:]))
	}
	return sums
}

func main() {
	total := Sum([]int{1, 2, 3})
	fmt.Println("Hello Sum", total)
}
