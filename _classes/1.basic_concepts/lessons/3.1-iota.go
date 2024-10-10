package lessons

import "fmt"

/*
 - Iota demonstra o uso do identificador iota em Go.
 - Iota é um identificador predefinido usado em declarações de constantes para simplificar definições de números incrementais.
 - Ele é redefinido para 0 sempre que a palavra-chave 'const' aparece no código-fonte e incrementa em 1 após cada especificação de constante.
 - Isso é particularmente útil para criar constantes enumeradas.
*/

const (
	// iota pode ser colocado em uma expressão
	iA = (iota + 1) * 2
	// a atribuição do iota não precisa se repetir
	iB
	iC
	// Valores podem ser pulados usando o _
	_
	iE
)

func Iota() {
	fmt.Println(iA, iB, iC, iE)
}
