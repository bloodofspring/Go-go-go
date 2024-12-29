package kata

import (
	"fmt"
	"strings"
)

func ListToRange(list []int) string {
	res := ""

	for i := 0; i < len(list); i++ {
		if i+1 < len(list) && list[i]+1 == list[i+1] {
			cur := i
			for cur+1 < len(list) && list[cur]+1 == list[cur+1] {
				cur++
			}

			if cur != i+1 {
				res += fmt.Sprintf("%d-%d,", list[i], list[cur])
				i = cur
				continue
			}
		}

		res += fmt.Sprintf("%d,", list[i])
	}

	return strings.TrimRight(res, ",")
}
