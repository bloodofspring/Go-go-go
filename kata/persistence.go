package kata

func Persistence(n int) int {
	if n < 10 {
		return 0
	}

	product := 1
	for n > 0 {
		product *= n % 10
		n /= 10
	}

	return 1 + Persistence(product)
}
