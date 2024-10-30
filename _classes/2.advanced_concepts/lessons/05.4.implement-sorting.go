package lessons

import (
	"fmt"
	"sort"
)

/*
	Com as interfaces é possível implementar customização em qualquer tipo de dado, seja ele da standard library ou não
	Para isso, basta implementar os métodos necessários para a interface desejada

	Para o exemplo, será implementado um sort de avaliação de livros
	Para implementar o sort, é necessário implementar os métodos Len, Less e Swap (https://pkg.go.dev/sort#Interface) em uma coleção de livros
*/

type Book struct {
	Name string
	Rate int
}

type BookByRate []Book

func (f BookByRate) Len() int {
	return len(f)
}

func (f BookByRate) Less(i, j int) bool {
	return f[i].Rate < f[j].Rate
}

func (f BookByRate) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func ImplementInterfaceSorting() {
	books := []Book{
		{"Livro 1", 4},
		{"Livro 2", 6},
		{"Livro 3", 1},
		{"Livro 4", 10},
		{"Livro 5", 0},
	}

	fmt.Printf("books unsorted: %v\n", books)

	sort.Sort(BookByRate(books))

	fmt.Printf("books sorted: %v\n", books)
}
