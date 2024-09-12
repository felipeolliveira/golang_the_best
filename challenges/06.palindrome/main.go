package main

import (
	"fmt"
	"strings"
)

func main() {

	playing := true

	for playing {
		fmt.Printf("\n\n======================\n")
		fmt.Printf("== É um palindromo? ==\n")
		fmt.Printf("======================\n")

		fmt.Printf("Digite uma palavra, frase ou uma sequencia números:\n")
		input := scanUserInput()
		input, err := sanitizeString(input)

		if err != nil {
			fmt.Printf("Ocorreu um erro ao checar a frase: %v\n", err)
		}

		response := checkPalindrome(input)

		if response {
			fmt.Printf("Sim, é um palíndromo!\n")
		} else {
			fmt.Printf("Não, não é um palíndromo!\n")
		}

		fmt.Printf("======================\n")
		fmt.Printf("Você gostaria de tentar outra palavra? (Y/n): ")
		input = strings.ToLower(scanUserInput())

		if input == "n" || input != "" {
			break
		}

		continue
	}
}
