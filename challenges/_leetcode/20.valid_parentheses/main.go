// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Example 1:
// Input: s = "()"
// Output: true

// Example 2:
// Input: s = "()[]{}"
// Output: true

// Example 3:
// Input: s = "(]"
// Output: false

// Example 4:
// Input: s = "([])"
// Output: true

// Constraints:
// 1 <= s.length <= 104
// s consists of parentheses only '()[]{}'.

package main

func main() {
	// r := isValid("({})[]")
	r := isValid("()))")
	println(r)
}

func isValid(s string) bool {
	stack := []rune{}

	if len(s)%2 == 1 {
		return false
	}

	for _, char := range s {

		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		lastOpenBracket := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		pair := string(lastOpenBracket) + string(char)

		if pair == "()" || pair == "[]" || pair == "{}" {
			continue
		}

		return false
	}

	return len(stack) == 0
}
