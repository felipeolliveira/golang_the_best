package main

import (
	"fmt"
)

func main() {
	r := addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011")
	// r := addBinary("1011", "1011") // 10110

	fmt.Printf("%s\n", r)
}

func addBinary(a string, b string) string {
	if len(b) > len(a) {
		return addBinary(b, a)
	}

	result := ""
	carry := 0

	for i := range a {
		// 48 is uint32 for rune "0"
		// 49 is uint32 for rune "1"
		// decrease 48, we have integer to both values, 0(48-48) and 1(49-48)
		valueA := int(a[len(a)-1-i]) - 48
		valueB := 0

		if i <= len(b)-1 {
			valueB = int(b[len(b)-1-i]) - 48
		}

		switch carry + valueA + valueB {
		case 0:
			result = "0" + result
			carry = 0
		case 1:
			result = "1" + result
			carry = 0
		case 2:
			result = "0" + result
			carry = 1
		case 3:
			result = "1" + result
			carry = 1
		}
	}

	if carry == 1 {
		result = "1" + result
	}

	return result
}
