package lessons

import "fmt"

func Pointers() {
	/*
		- Ponteiros são variáveis que armazenam o endereço de memória de outra variável
			p = 0xc0000b4008
			p -> 0xc0000b4008 -> ???(valor)
		- Ponteiros são usados para acessar variáveis indiretamente
		- A forma mais simples de ponteiros são de ponteiros de Stack (Stack Pointers)

		- Declaração de ponteiros:
			&variavel -> retorna o endereço de memória da variável
			*variavel -> retorna o valor da variável que o ponteiro aponta, ou acessa o valor da variável que o ponteiro aponta
	*/

	x := 10
	pointer := &x

	fmt.Println("poiter", pointer, *pointer)

	initialValue := 10
	changeValue(&initialValue)

	pointerByReturn := returnPointer()
	fmt.Println("newPointer", pointerByReturn, *pointerByReturn)

	/*
		CUIDADO COM ISSO
		- Quando a função que manipula um ponteiro recebe uma variavel por referencia não inicializada, o valor da variável é nil
			Caso a função tente fazer a dereferencia (*<var>), o programa irá quebrar com o erro:
			"panic: runtime error: invalid memory address or nil pointer dereference"
		- Go ainda não tem proteção contra execução de código com ponteiros inválidos (nil)
	*/
	var initialValuePointer *int
	changeValue(initialValuePointer) // Nesse momento, initialValuePointer é ponteiro com valor nil
}

/*
  - Toda vez que for usar um ponteiro, é necessário
    Colocar o * antes do tipo da variável
    Colocar o & antes da variável que deseja pegar o endereço no momento de passar a referencia
    `changeValue(&initialValue) -> func changeValue(x *int)`

  - A definição de um ponteiro com * pode ser ambiguo:
    *<type> -> ponteiro -> *int, *string, *float64, etc
    *<var> -> dereferencia -> alterar o valor da variável que o ponteiro aponta
*/
func changeValue(x *int) {
	*x = 50
}

/*
- Função que retorna um ponteiro é uma forma de retornar um valor de uma função
- A função retorna o endereço de memória da variável, e não o valor da variável
- A variável que recebe o retorno da função deve ser do mesmo tipo do ponteiro

- Não é possível fazer isso em C, nem em Rust, causaria um erro de compilação "undefined behavior / dangling pointer"
- Em Go não é possível fazer "pointer arithmetic" (aritmética de ponteiros), o que evita erros de memória e vazamentos de memória
*/
func returnPointer() *int {
	x := 10
	fmt.Println("returnPointer() ->", &x, x)
	return &x
}
