package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

var err error

const toZ int32 = 96

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func hash(str string) int {
	totalHash := 0
	for n, s := range str {
		totalHash += int(s-toZ) * int(math.Pow(37, float64(n)))
	}

	return totalHash
}

func HashTaskMain() {
	var str string

	_, _ = fmt.Scanf("%d")
	handleError(err)
	_, err = fmt.Scanf("%s", &str)
	handleError(err)

	fmt.Println(hash(str))
}

func AllDevours() {
	var n int
	_, err = fmt.Scanf("%d", &n)
	handleError(err)

	for i := 1; i <= n/2+1; i++ {
		if n%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("%d", n)
}

func Matrix() {
	var n int
	_, err = fmt.Scanf("%d", &n)
	handleError(err)

	graph := make(map[int][]string)
	reader := bufio.NewReader(os.Stdin)

	for i := 1; i <= n; i++ {
		var data string
		data, err = reader.ReadString('\n')
		handleError(err)

		graph[i] = strings.Split(strings.Trim(data, "\n"), " ")[1:]
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				fmt.Print("0 ")
			} else if slices.Contains(graph[i], strconv.Itoa(j)) {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Print("\n")
	}
}
