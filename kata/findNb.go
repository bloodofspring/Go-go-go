package kata

func FindNb(m int) int {
	n := 0
	for m > 0 {
		m -= n * n * n
		n++
	}

	if m != 0 {
		return -1
	}

	return n - 1
}
