package lessons

import "fmt"

/*
	# Any e Type Assertions

	- Há uma forma de driblar a tipagem de Go usando `any`, com isso, qualquer tipo pode ser atribuido a uma variável ou ser recebida por uma função
	- Usar com cautela pois pode ser perigoso e não é recomendado usar sem saber o que está fazendo
	- O ideal é que o uso do `any` sempre seja acompanhado de `type assertions` para garantir que o tipo seja o esperado
	- Para verificar o tipo de uma variavel:
		`result, ok := value.(<type>)`, sendo que o `ok` é um booleano que indica se o tipo é o esperado, e `result` é o valor convertido
*/

func isString(value any) bool {
	_, ok := value.(string)
	return ok
}

/*
	- É possível usar `type assertions` em conjunto com `switch` para verificar o tipo de uma variável
		No caso, deve se usar a palavra reservada `type` para pegar o tipo da variável
		value.(type)
*/

func checkType(value any) {
	switch t := value.(type) {
	case string:
		fmt.Println(t, "Value is a string")
	case int:
		fmt.Println(t, "Value is an integer")
	default:
		fmt.Println(t, "Value is another type")
	}
}

func InterfaceAndAssertions() {
	var value any = "Hello World"

	if isString(value) {
		println("Value is a string")
	} else {
		println("Value is not a string")
	}

	checkType(value)
}
