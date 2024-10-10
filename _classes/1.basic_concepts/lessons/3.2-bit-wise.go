package lessons

import "fmt"

// O operador bitwise para deslocar bits (>> | <<)
func BitWise() {
	x := 4      // 100
	y := x >> 1 // 10 (2)
	z := x << 1 // 1000 (8)
	fmt.Printf("x => %b\n", x)
	fmt.Printf("x >> 1 =>%b\n", y)
	fmt.Printf("x << 1 =>%b\n", z)

	fmt.Printf("================\n")
	fmt.Printf("KB, MB, GB => %v, %v, %v\n", KB, MB, GB)
}

// Exemplo de uso para criação de uma lista de numeros inteiros representando a quantidade de bytes por nomenclatura:
const (
	_  = iota
	KB = 1 << (iota * 10) // 1 << 10
	MB                    // 1 << 20
	GB                    // 1 << 30
)
