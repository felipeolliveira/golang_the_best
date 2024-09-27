package main

func romanToInt(s string) int {
	result := 0
	roman_characters := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i, char := range s {
		nextValue := 0
		value := roman_characters[string(char)]

		if i+1 < len(s) {
			nextValue = roman_characters[string(s[i+1])]
		}

		if nextValue > value {
			result -= value
		} else {
			result += value
		}
	}

	return result
}
