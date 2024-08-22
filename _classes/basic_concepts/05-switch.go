package basic_concepts

import (
	"fmt"
	"math"
)

func Switch() {
	/*
		Controle de fluxo com Switch bem comum em várias linguagens
	*/

	/*
		- Não precisa do `break` para parar a execução do switch
		- Não precisa de um case para cada valor, ele para na primeira condição verdadeira
		- Pode ter mais de um valor dentro do case
		- O Go não para a execução do switch após um erro de range
			Basta usar o default para tratar o caso não abordado
	*/
	numA := 2
	switch numA {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4, 5, 6:
		fmt.Println("three")
	default:
		fmt.Println("unknown number")
	}

	/*
		- Pode ser usado sem condição inicial, como se estivesse usando um if/else
	*/
	switch { // sem condição
	case numA == 1:
		fmt.Println("one")
	case numA == 2:
		fmt.Println("two")
	case "qualquer coisa" == "coisa qualquer": // um case totalmente aleatório
		fmt.Println("não")
	default:
		fmt.Println("default")
	}

	/*
		- Pode utilizar variaveis dentro do statement do switch
			No caso, a variavel `x` é declarada e atribuida com o valor de `math.Sqrt(4)`
			E o switch, assim como do if/else, precisa ter 2 statements
			`switch <variavel>; <condição> {}`
	*/
	switch x := math.Sqrt(4); x {
	case 2:
		fmt.Println("x is 2")
	default:
		fmt.Println("x is not 2")
	}

	/*
		o Switch também pode resolver casos de tipos
	*/
	whatAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("string")

}
