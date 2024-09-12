package main

import (
	"bufio"
	"os"
)

// scanUserInput reads a line of text from standard input and returns it as a string.
// It uses a bufio.Scanner to read the input from os.Stdin.
func scanUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}
