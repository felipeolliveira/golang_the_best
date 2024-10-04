package lessons

import "fmt"

/*
	- Tipos
	É possível criar os proprios tipos, se baseando em tipos primitivos ou compostos

	- Conversão
	Só é possível converter tipos equivalentes
*/

type custom int

var a int = 10
var b custom = 10

func CustomTypes() {
	fmt.Printf("a => %v, %T\n", a, a)
	fmt.Printf("b => %v, %T\n", b, b)

	// qualquer operação não pode ser realizar usando tipos diferentes
	// fmt.Printf("%v", a + b)

	// Convertendo o valor usando a sintaxe T(x)
	c := custom(a)
	fmt.Printf("c (conversão do a) => %v, %T\n", c, c)
	fmt.Printf("b + c => %v, %T\n", b+c, b+c)
}
