package main

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// sanitizeString takes an input string and returns a sanitized version of it.
//
// If an error occurs during the transformation process, it returns an empty string and the error.
//
// Parameters:
// - input: The string to be sanitized.
//
// Returns:
//   - A sanitized string with only lowercase alphanumeric characters.
//     For example, "Hêlló, World!" becomes "helloworld".
//   - An error if the transformation process fails.
func sanitizeString(input string) (string, error) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, err := transform.String(t, input)

	if err != nil {
		return "", err
	}

	regexp := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	result = regexp.ReplaceAllString(result, "")
	result = strings.ToLower(result)

	return result, nil
}
