package main

import "strings"

func main() {
	r := longestCommonPrefix([]string{
		// "flower", "flow", "flight",
		"dog", "racecar", "car",
	})
	println(r)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := ""
	for i, char := range strs[0] {
		prefix = prefix + string(char)

		for _, word := range strs {
			if !strings.HasPrefix(word, prefix) {
				return prefix[:i]
			}
		}
	}

	return prefix
}
