package main

import (
	"fmt"
	"main/kata"
)

func main() {
	// fmt.Println(kata.ListToRange([]int{1, 2, 3, 4, 8, 10, 12, 13, 14}))
	fmt.Println(kata.ListToRange([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))
	fmt.Println(kata.ListToRange([]int{40, 44, 48, 51, 52, 54, 55, 58, 67, 73}))
}
