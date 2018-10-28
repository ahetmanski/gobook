// Package `word` privides utilities for word's games
package word

// IsPalindrome returns true for palindrome string.
// (First try.)
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
