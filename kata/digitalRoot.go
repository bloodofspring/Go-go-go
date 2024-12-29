package kata

func DigitalRoot(n int) int {
	if n < 10 {
		return n
	}
	return DigitalRoot(n%10 + DigitalRoot(n/10))
}
