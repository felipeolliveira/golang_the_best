package lessons

import "fmt"

func Arrays() {
	/*
		- Arrays tem tamanho fixo em memória, mas Slice pode ser dinâmico e fixo.

		- Slice nada mais é que um ponteiro para um array que apenas mostra o valor de uma fatia (slice).
			arr[<min>:<max>] -> min é o índice inicial (inclusivo) e max é o índice final (não inclusivo).
			Esses limites que delimitam o slice não são obrigatórios pois inicial com o valor 0 e final com o tamanho do array `len(arr)`.

		- Slices tem lenth e capacity.
			Length: é o tamanho do slice
			Capacity: é o tamanho do array subjacente a partir do índice inicial do slice até o tamanho total sem fazer uma nova alocação de memória.

		- Apenas slices podem usar a função append() para adicionar novos elementos.
	*/

	arr := [5]int{1, 2, 3, 4, 5} // [1 2 3 4 5]
	fmt.Println("array", arr)

	slice := arr[1:3]           // [2 3]
	sliceWithNoMin := arr[:5]   // [1 2 3 4 5]
	sliceWithNoMax := arr[0:]   // [1 2 3 4 5]
	sliceWithNoMinMax := arr[:] // [1 2 3 4 5]
	sliceLength := len(slice)   // 2
	sliceCapacity := cap(slice) // 5

	fmt.Println("slice:", slice)
	fmt.Println("sliceWithNoMin:", sliceWithNoMin)
	fmt.Println("sliceWithNoMax:", sliceWithNoMax)
	fmt.Println("sliceWithNoMinMax:", sliceWithNoMinMax)
	fmt.Println("sliceLength, sliceCapacity:", sliceLength, sliceCapacity)

	/*
		- Se alterarmos o valor de um elemento do slice, o valor do elemento correspondente no array também será alterado.
	*/
	slice[0] = 100
	fmt.Println("Array", arr)   // [1 100 3 4 5]
	fmt.Println("Slice", slice) // [100 3]

	/*
		- Como dito antes, o Slice pode ser iniciado omitindo o tamanho do array.
	*/
	initalSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("initalSlice", initalSlice) // [100 3]

	getMoviesByUser()

	/*
		- É possível criar um slice de slices, matrizes.
	*/
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("matrix", matrix) // [[1 2 3] [4 5 6]]
}

func getMoviesByUser() {
	moviesInDB := []string{"De Volta ro Futuro", "De Volta ro Futuro 2", "Interestelar", "Madagascar", "Jurassic Park", "Jurassic Park 2", "Jurassic Park 3"}
	moviesIds := []int{1, 2, 3, 4, 5}
	fmt.Println("moviesIds", len(moviesIds), cap(moviesIds))

	/*
		- Slice dinâmico por não saber o tamanho do array subjacente.
			Não pode ser declarado com short syntax :=
	*/
	var dynamicSliceForMovieTitle []string

	// A cada iteração, o slice é realocado para um novo array com o dobro de capacidade.
	// 0, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512...
	fmt.Println("Dynamic Slice: Init ==========")
	for _, movieId := range moviesIds {
		movie := moviesInDB[movieId]
		fmt.Println("len:", len(dynamicSliceForMovieTitle), "cap:", cap(dynamicSliceForMovieTitle))
		dynamicSliceForMovieTitle = append(dynamicSliceForMovieTitle, movie)
		fmt.Println("len:", len(dynamicSliceForMovieTitle), "cap:", cap(dynamicSliceForMovieTitle))
	}
	fmt.Println("Dynamic Slice: End ==========")

	/*
		- Slice fixo por saber o tamanho da requisição e dos dados.
		Pode ser declarado com short syntax :=
		utiliza make([]T, len, cap) para definir o tamanho do slice.
		É recomendado usar quando se sabe o tamanho do array subjacente, pois não precisa realocar memória a cada iteração.
	*/
	staticSliceForMovieTitle := make([]string, 0, 5)

	fmt.Println("Static Slice: Init ==========")
	// A cada iteração, o slice é realocado para um novo array com o dobro de capacidade.
	// 0, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512...
	for _, movieId := range moviesIds {
		movie := moviesInDB[movieId]
		fmt.Println("len:", len(staticSliceForMovieTitle), "cap:", cap(staticSliceForMovieTitle))
		staticSliceForMovieTitle = append(staticSliceForMovieTitle, movie)
		fmt.Println("len:", len(staticSliceForMovieTitle), "cap:", cap(staticSliceForMovieTitle))
	}
	fmt.Println("Static Slice: End ==========")
}
