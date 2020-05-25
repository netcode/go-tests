package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	expected := 15

	if got != expected {
		t.Errorf("Got %d, expected %d : Given %v", got, expected, numbers)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("Test SumAll for multiple arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	})

	t.Run("Test SumAll for single array", func(t *testing.T) {
		got := SumAll([]int{3, 2, 3})
		expected := []int{8}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	})

}
