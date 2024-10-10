package main

import "fmt"

func main() {
	r := resursive(15)
	fmt.Println(r)

	r = iterative(15)
	fmt.Println(r)
}

func resursive(n int) int {
	if n <= 1 {
		return 1
	}

	return n * resursive(n-1)
}

func iterative(n int) int {
	result := 1

	for n > 1 {
		result *= n
		n--
	}

	return result
}
