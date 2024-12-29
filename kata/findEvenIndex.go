package kata

func Sum(arr []int) int {
	res := 0
	for _, v := range arr {
		res += v
	}
	return res
}

func FindEvenIndex(arr []int) int {
	var left, right int

	for i := 0; i < len(arr); i++ {
		left = Sum(arr[:i])
		right = Sum(arr[i+1:])

		if left == right {
			return i
		}
	}

	return -1
}
