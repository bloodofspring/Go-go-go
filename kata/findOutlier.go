package kata

import "math"

func compare(a, b int) bool {
	return math.Abs(float64(a%2)) == math.Abs(float64(b%2))
}

func FindOutlier(integers []int) int {
	outlier := integers[0]
	if !compare(integers[1], outlier) {
		if !compare(integers[1], integers[2]) {
			return integers[1]
		}
		return outlier
	}

	for _, v := range integers {
		if !compare(v, outlier) {
			return v
		}
	}

	return -1
}
