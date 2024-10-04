package main

import "strings"

func main() {
	r := lengthOfLastWord("   fly me   to   the moon  ")
	println(r)
}

func lengthOfLastWord(s string) int {
	splitted := strings.Split(strings.Trim(s, " "), " ")

	return len(splitted[len(splitted)-1])
}
