package kata

import "strings"

func Disemvowel(comment string) string {
	var vowels = []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

	for _, vowel := range vowels {
		comment = strings.ReplaceAll(comment, vowel, "")
	}

	return comment
}
