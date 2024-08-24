package lessons

import "fmt"

func Arrays() {
	/*
		Arrays em Go são uma coleção de elementos do mesmo tipo.
		- O tamanho do array é parte do seu tipo, então arrays de tamanhos diferentes são tipos diferentes.
		- Arrays são zero-valued, ou seja, são inicializados com o valor zero do tipo.
		- Arrays são imutáveis, ou seja, não podem ser alterados depois de criados.
		- O Array pode ser iniciado com valores ou com valores em índices específicos.
	*/
	zeroValuedArray := [10]int{}
	filledArray := [5]int{1, 2, 3, 4, 5}
	filledByIndexArray := [10]int{4: 400, 7: 300} // O índice 4 recebe o valor 400, o índice 7 recebe o valor 300

	fmt.Println(zeroValuedArray)
	fmt.Println(filledArray)
	fmt.Println(filledByIndexArray)

	stringArray := [10]string{5: "apple", 7: "banana"}
	fmt.Println(stringArray)

	/*
		Arrays podem ter seu tamanho inferido pelo compilador, basta deixar o tamanho vazio.
		Esse tipo de dado é chamado de Slice.
	*/
	inferredArray := []int{1, 2, 3, 4, 5}
	fmt.Println(inferredArray)

	/*
		Arrays podem ter seu tamanho declarados apenas por contantes inteiras.
		- Variaveis não podem ser usadas para definir o tamanho de um array.
	*/
	const size = 5

	sizedArrayByConst := [size]int{1, 2, 3, 4, 5}
	fmt.Println(sizedArrayByConst)
}
