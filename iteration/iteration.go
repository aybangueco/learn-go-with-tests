package iteration

import "strings"

// Repeat a character based on the repeat count.
func Repeat(character string, repeatCount int) string {
	var repeated strings.Builder

	for range repeatCount {
		repeated.WriteString(character)
	}

	return repeated.String()
}
