package kata

import "regexp"

func GetCount(str string) (count int) {
	return len(regexp.MustCompile("[aeiou]").FindAllString(str, -1))
}
