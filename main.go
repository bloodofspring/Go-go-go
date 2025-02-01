package main

import (
	"fmt"
	"main/useful_things"
)

// func printSnake(snake [][]int) {
// 	for _, row := range snake {
// 		for i := 0; i < len(row); i++ {
// 			if row[i] == 1 {
// 				fmt.Print("@")
// 			} else {
// 				fmt.Print("-")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

func main() {
	fmt.Println(useful_things.RandomDate())
}
