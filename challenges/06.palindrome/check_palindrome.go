package main

func checkPalindrome(input string) bool {
	for i := range len(input) / 2 {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
