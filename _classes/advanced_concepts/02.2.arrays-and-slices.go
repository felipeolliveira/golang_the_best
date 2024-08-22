package advanced_concepts

import "fmt"

func ArraysAndSlices() {
	/*
		- Slices com 3 indices podem manipular o tamanho e a capacidade do array subjacente.
			O primeiro índice é o índice inicial do slice.
			O segundo índice é o índice final do slice.
			O terceiro índice é a capacidade do slice.
			arr[0:2:2] -> slice com índices 0 e 1 e capacidade 2.

			Não é possível exceder a capacidade do array de onde foi feito o slice.
			Não é possível omitir o primeiro índice e informar o terceiro.
	*/
	arr := [5]int{1, 2, 3, 4, 5}
	sliceWithNoCapIndex := arr[0:2]
	sliceWithCapIndex := arr[0:2:2]

	fmt.Println(sliceWithNoCapIndex, cap(sliceWithNoCapIndex))
	fmt.Println(sliceWithCapIndex, cap(sliceWithCapIndex))

	/*
		- Um valor de um slice só pode ser acessado com um indíce válido.
		Toda vez que o Go tenta acessar um índice é realizado um bound checking. O Bound checking é feito em tempo de execução e não compilação.
		Funciona como uma proteção para acessar índices inválidos.

		- É possível "ajudar" o compilador e não fazer o bound checking a todo momento, fazendo um "hint" para o compilador.
		Para isso, basta atribuir a uma variavel _ o indice máximo que será acessado, pois o compilador saberá que o indice informado é o máximo,
		e não fará o bound checking para índices menores que o informado.
		Exemplo: acesse a biblioteca bynary.BigEndian.PutUint64 e veja como é feito o bound checking.
	*/
	sliceBoundChecking := arr[:]
	valid1 := sliceBoundChecking[0] // 1 bound checking
	valid2 := sliceBoundChecking[1] // 2 bound checking
	valid3 := sliceBoundChecking[2] // 3 bound checking
	// invalid := sliceBoundChecking[10] // panic: runtime error: index out of range [3] with length 2
	fmt.Println("sliceBoundChecking", valid1, valid2, valid3)

	// go run -gcflags="-d=ssa/check_bce" . para ver o bound checking
	sliceOptimezedBoundChecking := arr[:]
	_ = sliceOptimezedBoundChecking[4]      // apenas esse bound checking será feito
	valid1 = sliceOptimezedBoundChecking[1] // ...
	valid2 = sliceOptimezedBoundChecking[2] // ...
	valid3 = sliceOptimezedBoundChecking[3] // ...

	fmt.Println("sliceOptimezedBoundChecking", valid1, valid2, valid3)

	/*
		- Slices são passados para funções por ponteiros por padrão. Os arrays são passados por valor por padrão.
	*/
	arrayToChange := [3]int{1, 2, 3}
	sliceToChange := []int{1, 2, 3}

	changeArray(arrayToChange)
	changeSlice(sliceToChange)

	fmt.Println(arrayToChange) // [1 2 3]
	fmt.Println(sliceToChange) // [100 2 3]
}

func changeSlice(slice []int) {
	slice[0] = 100
}

func changeArray(array [3]int) [3]int {
	array[0] = 100
	return array
}
