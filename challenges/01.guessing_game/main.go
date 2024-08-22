package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Introduction to the game
	const spacerln = "====================================="

	fmt.Println(spacerln)
	fmt.Println("Jogo de Adivinhação")
	fmt.Println("Um número aleatório será sorteado. Tente acertar! O número é um inteiro entre 1 e 100.")
	fmt.Println("[Digite um número e pressione Enter. Você tem 10 tentativas]")
	fmt.Println(spacerln)

	randomInt := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	attempts := [10]int64{}

	for attemptIndex := range attempts {
		fmt.Printf("[%d/10]: Qual o seu chute?\n", attemptIndex+1)
		fmt.Printf("➡➡➡ ")
		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)

		inputInt, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println("⛔ | O seu chute precisa ter um número inteiro de 0 até 100! | ⛔")
			return
		}

		switch {
		case inputInt < randomInt:
			fmt.Printf("Você error. O número sorteado é maior que %d ⬆️\n", inputInt)
		case inputInt > randomInt:
			fmt.Printf("Você error. O número sorteado é menor que %d ⬇️\n", inputInt)
		case inputInt == randomInt:
			fmt.Printf(spacerln + "\n" + spacerln + "\n\n")
			fmt.Printf(
				"🎉🎉 Você acertou. É o número %d! 🎉🎉\n"+
					"Essas foram suas tentativas: %v\n",
				randomInt,
				attempts[:attemptIndex],
			)
			fmt.Printf("\n" + spacerln + "\n" + spacerln + "\n")
			return
		}

		fmt.Println(spacerln)

		attempts[attemptIndex] = inputInt
	}

	fmt.Printf(spacerln + "\n" + spacerln + "\n\n")
	fmt.Printf(
		"😢 Infelizmente você não acertou.O número era: %d.\n"+
			"Essas foram suas tentativas: %v\n",
		randomInt,
		attempts,
	)
	fmt.Printf("\n" + spacerln + "\n" + spacerln + "\n")
}
