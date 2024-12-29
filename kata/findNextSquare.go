package kata

import "math"

func FindNextSquare(sq int64) int64 {
	if sq%int64(math.Sqrt(float64(sq))) != 0 {
		return -1
	}
	return int64(math.Pow(math.Sqrt(float64(sq))+1, 2))
}
