package kata

func Spiralize(size int) [][]int {
    // Make a snake
	var totalSnake = [][]int{}

	for i := 0; i < size; i++ {
		row := []int{}
		for i := 0; i < size; i++ {
			row = append(row, 0)
		}
		totalSnake = append(totalSnake, row)
	}

	

	return totalSnake
}