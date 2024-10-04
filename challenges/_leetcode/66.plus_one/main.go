package main

import (
	"fmt"
)

func main() {
	r := plusOne([]int{
		// 9, 9, 9, 9,
		4, 9, 8, 9,
		// 5, 0, 0, 0,
	})

	fmt.Printf("%v\n", r)
}

func plusOne(digits []int) []int {
	result := digits

	for i := len(result) - 1; i >= 0; i-- {
		if result[i] < 9 {
			result[i]++
			return result
		}
		result[i] = 0
	}

	result = append([]int{1}, result...)
	return result
}
