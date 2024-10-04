package lessons

import "fmt"

/*
Valores primitivos no Go são os valores básicos que são usados para representar dados.
Semelhantes a outras linguagens de programação, Go possui tipos de dados primitivos como inteiros, flutuantes, booleanos, strings, etc.
*/

func Primitives() {
	// Strings
	// Podem ser concatenadas com o operador `+`
	fmt.Println("go" + "lang")

	// Strings podem ser definidas como literal ou raw
	// - Literal: será interpretado
	// - Raw: não será interpretado, mostrando o texto cru
	fmt.Println("is literal\n is\t literal")
	fmt.Println(`is literal\n is\t literal`)

	// Inteiros e Flutuantes
	// Podem ser manipulados com operadores aritméticos normalmente
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7/3.0)

	// Booleanos
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
