package lessons

import "fmt"

/*
Variáveis em Go tem dois estagios de uso: declaração e inicialização.
Ambos podem ser feitos em uma única linha ou separadamente.
*/

func Variables() {
	// As variaveis podem ser iniciadas com a palavra reservada var
	var initialString = "initial"
	fmt.Println(initialString)

	// Pode ser iniciado mais de uma var na mesma linha e com o mesmo tipo definido
	var intA, intB int = 1, 2
	// Ou com tipos diferentes, porém não é possível definir o tipo já que eles são inferidos
	var stringC, floatD = "string", 1.0
	fmt.Println(intA, intB, stringC, floatD)

	// Pode ser iniciado mais de uma var no mesmo escopo do `var()` e com tipos diferentes
	var (
		name     = "John"
		lastName = "Doe"
		age      = 25
	)
	fmt.Println(name, lastName, age)

	var boolean = true
	fmt.Println(boolean)

	/*
		Go atribui o valor um conceito de `zero-valued` para variaveis declaradas mas não inicializadas

		- int: 0
			Por padrão, inteiros são inteiros de 32bits, mas podem ser alterados para outros tamanhos
			como int8, int16, int32, int64
		- uint: 0
			Por padrão, inteiros sem sinal são uinteiros de 32bits, mas podem ser alterados para outros tamanhos
			como uint8, uint16, uint32, uint64, uintptr, byte(alias para uint8), rune(alias para int32)
		- float: 0.0
			Por padrão, floats são float64, mas podem ser alterados para float32
		- string: ""
		- boolean: false
		- complexos(ponteiros, functions): nil
	*/
	var e int
	fmt.Println(e)

	// A declaração de variaveis pode ser feita de forma simplificada
	// O tipo da variavel é inferido pelo valor atribuido
	// A declaração de variaveis simplificada só pode ser feita dentro de funções
	appleString := "apple"
	fmt.Println(appleString)

	/*
		As variaveis podem ser privadas ou publicas, dependendo da primeira letra do nome da variavel
		- Variaveis com a primeira letra maiuscula são publicas
		- Variaveis com a primeira letra minuscula são privadas
	*/
	var PublicString = "Public"
	var privateString = "private"
	fmt.Println(PublicString, privateString)
}
