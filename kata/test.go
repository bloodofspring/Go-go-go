package kata

import "slices"

func diameterSlice(m [][]int, border int) []int {
	var result []int

	result = append(result, m[border][border:len(m)-border]...)

	if border == (len(m)-1)/2 && len(m)%2 != 0 {
		return result
	}

	for i := border + 1; i < len(m)-border-1; i++ {
		result = append(result, m[i][len(m)-border-1])
	}

	bottomRow := m[len(m)-border-1][border : len(m)-border]
	slices.Reverse(bottomRow)
	// bottomRow = reverseSlice(bottomRow)
	result = append(result, bottomRow...)

	for i := len(m) - border - 2; i > border; i-- {
		result = append(result, m[i][border])
	}

	return result
}

func Snail(snaipMap [][]int) []int {
	var result []int

	if len(snaipMap) == 0 || len(snaipMap[0]) == 0 {
		return []int{}
	}

	for i := 0; i < (len(snaipMap)+1)/2; i++ {
		result = append(result, diameterSlice(snaipMap, i)...)
	}

	return result
}

// for go version <= 1.20
// func reverseSlice(s []int) []int {
// 	result := []int{}

// 	for i := len(s) - 1; i >= 0; i-- {
// 		result = append(result, s[i])
// 	}

// 	return result
// }
