package main

import (
	"bufio"
	"os"
)

func scanUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}
