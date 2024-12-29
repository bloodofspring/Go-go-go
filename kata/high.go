package kata

import (
	"strings"
)

func High(s string) string {
	maxScore := 0
	maxWord := ""

	for _, word := range strings.Fields(s) {
		score := 0
		for _, c := range word {
			score += int(c - 'a' + 1)
		}

		if score > maxScore {
			maxScore = score
			maxWord = word
		}
	}

	return maxWord
}
