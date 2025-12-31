package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

var currencyRates = map[string]float64{
	"USD": 0.151,
	"EUR": 0.137,
	"JPY": 16.29,
	"GBP": 0.13,
	"CHF": 0.1402,
	"AUD": 0.2712,
	"CAD": 0.2374,
	"CNY": 1.251,
	"HKD": 1.326,
	"NZD": 0.2922,
	"SEK": 1.655,
	"NOK": 1.806,
	"DKK": 1.122,
	"SGD": 0.2249,
	"KRW": 242.97,
	"ZAR": 3.239,
	"MXN": 3.454,
	"INR": 14.71,
	"ILS": 0.63,
	"THB": 5.74,
	"IDR": 2875.0,
	"MYR": 0.754,
	"PHP": 9.74,
	"PLN": 0.644,
	"CZK": 3.77,
	"HUF": 61.59,
	"TRY": 6.49,
	"BGN": 0.293,
	"RON": 0.746,
}

var currencyKeys = slices.Collect(maps.Keys(currencyRates))

func main() {
	fmt.Println("Welcome to BRL Currency Converter!")
	fmt.Println("Processing values...")

	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Printf("===================\n")
		fmt.Printf("Format Error. Use <value> <currency> to convert.")
		fmt.Printf("===================\n")
		os.Exit(1)
	}

	value, err := strconv.ParseFloat(args[0], 64)

	if err != nil {
		fmt.Printf("===================\n")
		fmt.Printf("Value must be a number\n")
		fmt.Printf("===================\n")
		os.Exit(1)
	}

	rate, ok := currencyRates[strings.ToUpper(args[1])]

	if !ok {
		fmt.Printf("===================\n")
		fmt.Printf("Currency not supported. Try with one of the following:\n [%s]\n", strings.Join(currencyKeys, ", "))
		fmt.Printf("===================\n")
		os.Exit(1)
	}

	fmt.Printf("===================\n")
	fmt.Printf("BRL => %s: %.2f\n", args[1], value*rate)
	fmt.Printf("===================\n")
}
