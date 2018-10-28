// Package `word` provides utilities for word's games.
package word

import "unicode"

// IsPalindrome returns true for palindrome string.
// Ignore symbols registry and symbols, that are not letters.
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
