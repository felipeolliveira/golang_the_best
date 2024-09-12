package main

// checkPalindrome checks if the given input string is a palindrome.
// Parameters:
//
//	input (string): The string to be checked.
//
// Returns:
//
//	bool: Returns true if the input string is a palindrome, otherwise false.
func checkPalindrome(input string) bool {
	for i := range len(input) / 2 {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
