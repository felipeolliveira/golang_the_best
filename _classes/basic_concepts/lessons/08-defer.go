package lessons

import "fmt"

func Defer() {
	doDefer()
	doDeferLifo()
	doDeferImmediatelyEvaluated()
}

/*
  - Defer é um statement que adia a execução de uma função até o final do bloco da função atual (apenas funções)
    É comum ser usado para garantir que um recurso seja fechado, como um arquivo ou conexão de rede
*/
func doDefer() int {
	defer fmt.Println("world")
	fmt.Println("hello")
	return 1
}

/*
- As chamadas de defer são empilhadas (LIFO) e executadas na ordem inversa após o final do bloco
*/
func doDeferLifo() {
	defer fmt.Println("3")
	defer fmt.Println("2")
	defer fmt.Println("1")
	fmt.Println("counting...")
}

/*
  - O *argumento* da função é avaliado imediatamente (Immediately Evaluated), mas a chamada da função não é executada até o final do bloco
    Ou seja, o valor passado para o argumento da função `defer` é no momento que a função é chamada, caso o valor
    for alterado depois da chamada da função `defer`, o valor passado para a função `defer` não será alterado.
    Caso o valor dentro da função `defer` seja um closure(variavel de escopo acima) o valor será o valor atualizado
*/
func doDeferImmediatelyEvaluated() {
	x := 10
	defer func(y int) {
		fmt.Println("inside defer with arg:", y)
	}(x)

	defer func() {
		fmt.Println("inside defer without param (closure):", x)
	}()

	x = 50
	fmt.Println("outside defer", x)
	/*
		- Output:
			1. outside defer 50
			2. (LIFO) inside defer without param (closure): 50
			3. (LIFO) inside defer with arg: 10
	*/
}
