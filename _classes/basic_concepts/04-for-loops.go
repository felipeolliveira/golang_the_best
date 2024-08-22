package basic_concepts

import (
	"fmt"
)

func ForLoops() {

	/*
		ForLoops com declaração e condição. Bem parecido com ForLoops em outras linguagens.
	*/
	for j := 0; j <= 3; j++ {
		fmt.Println("common loop:", j)
	}

	/*
		ForLoops com apenas uma condição. Bem parecido com o while de outras linguagens.
	*/
	i := 0
	for i <= 3 {
		fmt.Println("for with only condition", i)
		// i = i + 1
		i++
	}

	/*
		Loops podem ser feitos com a palavra `range` para iterar sobre arrays, slices, maps, strings e canais.
		- Para arrays e slices, ele retorna o index e o valor.
		- Para strings, ele retorna o index e o valor do byte.
		- Para maps, ele retorna a chave e o valor.
		- caso você não queira usar o index ou a chave, você pode usar o `_` para ignorar.
	*/
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("using `range` with index:", index, value)
	}
	for _, value := range arr {
		fmt.Println("using `range` without index:", value)
	}

	/*
		ForLoops podem fazer o range apenas de inteiros, com isso simplifica o statement da iteração.
		- `range 6` é o mesmo que `range 0..6`.
		- Nesse caso, é possível declarar apenas uma variavel para o index, pois não existe valor
			como no caso de iteração por vetores, maps, slices, etc.

		ForLoops podem utilizar as palavras `break` e `continue` para controlar o fluxo.
		- `break` para a execução do loop.
		- `continue` para a execução atual e vai para a próxima.
	*/
	for n := range 6 {
		if n%2 == 0 {
			fmt.Println("next iteration in loop with the `continue`. Even:", n)
			continue
		}
		fmt.Println("Odd:", n)
	}
	for {
		fmt.Println("break a loop with word `break`")
		break
	}
}
