package go_strings

import (
	"strings"
	"unicode"
)

// Reverse returns a new string with the characters of the input string in reverse order.
// It uses a rune slice to ensure correct handling of Unicode characters.
//
// Parameters:
//   - s: the string to be reversed
//
// Returns:
//   - the reversed string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// StripNonAlphanumeric removes all characters from the input string except letters (Unicode),
// digits (Unicode), and spaces. It uses a strings.Builder for efficient string construction.
// This function ensures correct handling of Unicode characters by iterating over runes.
//
// Parameters:
//   - s: the input string to process
//
// Returns:
//   - a new string containing only alphanumeric characters and spaces from the input
func StripNonAlphanumeric(s string) string {
	var result strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == ' ' {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// StripBlanks removes all spaces (' ') from the input string.
// It iterates through each character of the string, and if a space is found,
// it is removed by slicing the string and concatenating the remaining parts.
// Note: This implementation may have performance issues for large strings or
// multiple consecutive spaces due to repeated string reallocations.
//
// Parameters:
//   - s: the input string to process
//
// Returns:
//   - a new string with all spaces removed
func StripBlanks(s string) string {
	var result strings.Builder
	for _, r := range s {
		if r != ' ' {
			result.WriteRune(r)
		}
	}
	return result.String()
}
