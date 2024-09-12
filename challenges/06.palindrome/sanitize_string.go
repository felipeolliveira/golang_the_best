package main

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

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
