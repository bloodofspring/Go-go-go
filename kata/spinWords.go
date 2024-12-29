package kata

import "strings"

func Reverse(s string) string {
	res := make([]string, len(s))
	for i, c := range s {
		res[len(s)-i-1] = string(c)
	}
	return strings.Join(res, "")
}

func SpinWords(str string) string {
	res := ""

	for _, word := range strings.Split(str, " ") {
		if len(word) >= 5 {
			res += Reverse(word) + " "
			continue
		}

		res += word + " "
	}

	return strings.TrimSpace(res)
}
