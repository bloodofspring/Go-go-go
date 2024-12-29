package kata

import (
	"strings"
)

func HasUniqueChar(str string) bool {
	for _, char := range str {
		if strings.Count(str, string(char)) != 1 {
			return false
		}
	}
	return true
}
