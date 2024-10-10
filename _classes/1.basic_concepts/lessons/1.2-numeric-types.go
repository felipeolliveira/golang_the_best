package lessons

import (
	"fmt"
	"runtime"
)

/*
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int					auto set int32 or int64 based on system arch
int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float					auto set float32 or float64 based on system arch
float32     the set of all IEEE 754 32-bit floating-point numbers
float64     the set of all IEEE 754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
*/

func NumericTypes() {
	// Cada rune do utf-8 pode consumir até 32 bits (4 bytes)
	a := "e"
	b := "é"
	c := "香"

	fmt.Printf("a, b, c => %v, %v, %v\n", a, b, c)
	fmt.Printf("bytes => %v, %v, %v\n", []byte(a), []byte(b), []byte(c))

	// O int e float são definidos como 32bits ou 64bits de acordo com a arquitetura do processador e OS
	fmt.Printf("OS => %v\n", runtime.GOOS)
	fmt.Printf("Arch => %v\n", runtime.GOARCH)
}
