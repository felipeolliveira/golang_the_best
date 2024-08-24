package lessons

import (
	"fmt"
)

func IfElse() {
	// Um controle de fluxo comum, nada demais...
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Controle de fluxo com if, else if e else também comum, nada demais...
	num := 10
	if num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "is less than 10")
	} else {
		fmt.Println(num, "is ten or more")
	}

	// A variável declarada no bloco if está disponível para o bloco else
	if num := 9; num > 0 {
		fmt.Println(num, "is positive, but variable is inside if block")
	}
}
