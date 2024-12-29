package kata

import "regexp"

func Alphanumeric(str string) bool {
	return regexp.MustCompile(`^[\w\d]+$`).MatchString(str)
}
