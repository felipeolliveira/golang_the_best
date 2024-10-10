package lessons

import (
	"fmt"
	"math"
)

/*
	Constantes em Go são declaradas com a palavra reservada `const`.
*/

/*
  - Uma constante pode ser de qualquer tipo de dado primitivo,
    mas é comum serem usadas para strings e números.
  - Constantes não podem ser declaradas usando short syntax `:=`.
  - Constantes não são obrigadas a serem usadas, o compilador não reclama.
  - Toda vez que uma constante tem o tipo declarado, o seu tipo é estrito e não pode ser usado nenhum outro tipo similar:
    Exemplo: `const pi float64 = 3.14` só pode ser usado em lugares que exijam float64.
*/
const typedStringConst string = "constante de string (typed)"

/*
  - Constantes ganha `untyped` quando o tipo é omitido e o Go infere a partir do contexto,
    isso significa que o tipo não é estrito e pode ser usado em qualquer lugar que aceite um tipo compatível,
    como um rune ou byte nesse caso.
    Exemplo: `const pi = 3.14` pode ser usado em lugares que exijam qualquer float, sem problemas.

- O tipo da contante é atribuido apenas no tempo de compilação
*/
const untypedStringConst = "constante de string (untyped)"

func Constants() {
	fmt.Println(typedStringConst)
	fmt.Println(untypedStringConst)

	/*
		Constantes numéricas recebem um tipo baseado no contexto,
		mas podem sofrer coerção de tipo, desde que o valor seja compatível.
		int -> float64
		float64 -> int
		...
	*/
	const huge = 500000000        // untyped int: Foi inferido como int
	const aritmetic = 3e20 / huge // untyped float64: Foi inferido como float64 depois da operação
	const hugeInt = int(huge)     // coerção de tipo, virou apenas int64

	fmt.Println(aritmetic)
	fmt.Println(hugeInt)

	// Nesse caso, a constante com tipo inferida (untyped) ou declarada podem ser usadas do mesmo jeito
	// No fim, a paravra `untyped` é apenas um detalhe de que o tipo foi inferido pelo Go
	fmt.Println(math.Sin(aritmetic))
}
