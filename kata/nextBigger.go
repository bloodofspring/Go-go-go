package kata

func NextBigger(n int) int {
	digits := getDigits(n)
	length := len(digits)

	i := length - 2
	for i >= 0 && digits[i] >= digits[i+1] {
		i--
	}

	if i < 0 {
		return -1
	}

	j := length - 1
	for j > i && digits[j] <= digits[i] {
		j--
	}

	digits[i], digits[j] = digits[j], digits[i]

	left := i + 1
	right := length - 1
	for left < right {
		digits[left], digits[right] = digits[right], digits[left]
		left++
		right--
	}

	result := 0
	for _, d := range digits {
		result = result*10 + d
	}

	return result
}

func getDigits(n int) []int {
	if n == 0 {
		return []int{0}
	}

	var digits []int
	for n > 0 {
		digits = append([]int{n % 10}, digits...)
		n /= 10
	}
	return digits
}
