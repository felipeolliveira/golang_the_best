package lessons

import (
	"fmt"
)

/*
rune        alias for int32
*/

func Utf8OrAscii() {
	s := "ascii éøâ 香 ❤"

	// O range transforma cada caracter em um `rune` e entende que a string é utf-8, ou seja:
	// - ascii: rune de int32, mas que usa apenas 1 byte que poderia ser int8
	// - éøâ: rune de int32, mas que usa 2 bytes dos 4 disponíveis
	// - 香: rune de int32, mas que usa 3 bytes dos 4 disponíveis
	for _, v := range s {
		fmt.Printf("%b - %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println("")

	// O For loop acessa cada byte, sendo que para os caracteres:
	//  - ascii: mostra completamente, pois usa apenas 1 byte, que corresponde a uma iteração do byte do loop
	// - éøâ: são necessários 2 bytes
	// - 香: são necessários 3 bytes
	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}
}
