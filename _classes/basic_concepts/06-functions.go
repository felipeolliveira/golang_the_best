package basic_concepts

import "fmt"

func Functions() {
	fmt.Println(sum(1, 2))                  // 1 + 2 => 3
	fmt.Println(swap(1, 2))                 // 1, 2 => 2, 1
	fmt.Println(divide(10, 3))              // 10 / 3 => 3, 10 % 3 => 1
	fmt.Println(hoc(10)(5))                 // 10 + 5 => 15
	fmt.Println(variadicSum(1, 2, 3, 4, 5)) // 1 + 2 + 3 + 4 + 5 => 15
}

/*
  - Pode ser omitido o tipo de retorno se todos os argumentos forem do mesmo tipo, colocando o tipo de retorno no final
    Exemplo: func sum(a, b int) int => func sum(a int, b int) int

- Go não tem sobrecarga de funções, então não é possível ter funções com o mesmo nome mas com argumentos diferentes
*/
func sum(a, b int) int {
	return a + b
}

/*
- Funções podem retornar mais de um valor
*/
func swap(a, b int) (int, int) {
	return b, a
}

/*
- Os retornos podem ser nomeados, o que é útil para documentar o que a função retorna
- Se os retornos forem nomeados, a declaração das variaveis internas da função não são necessárias
*/
func divide(a, b int) (resp int, rest int) {
	resp = a / b
	rest = a % b
	/*
		- Naked return: retorna os valores nomeados da função sem a necessidade de declarar os valores
			Apenas `return` é o mesmo que `return resp, rest`
			Não é muito recomendado, com excessão de funções bem pequenas. O motivo é pela ligibilidade do código
	*/
	return resp, rest
}

/*
  - HOC (Higher Order Function): Função que retorna outra função
    Dentro da função HOC, tem a implementação de uma `função anônima` que é retornada,
    no caso abaixo, a função anônima é também uma closure, pois ela tem acesso a variável `a` da função `hoc`
*/
func hoc(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

/*
  - Variadic Function: Função que aceita um número variável de argumentos
    O argumento variável deve ser o último argumento da função
*/
func variadicSum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
