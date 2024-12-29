package kata

import (
	"strings"
)

func ToCamelCase(s string) string {
	var output string

	if s == "" {
		return ""
	}

	for _, word := range strings.Split(s, "-") {
		if strings.Contains(word, "_") {
			for _, w := range strings.Split(word, "_") {
				output += strings.ToUpper(w[:1]) + w[1:]
			}
			continue
		}

		output += strings.ToUpper(word[:1]) + word[1:]
	}

	if s[:1] == strings.ToLower(s[:1]) {
		output = strings.ToLower(output[:1]) + output[1:]
	}

	return output
}
