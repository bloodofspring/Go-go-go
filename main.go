package main

import (
	"bufio"
	"fmt"
	"main/useful_things"
	"os"
)

func main() {
	var eq string

	fmt.Println("input equation:")
	reader := bufio.NewReader(os.Stdin)
	eq, err := reader.ReadString('\n')

	if err != nil {
		return
	}

	fmt.Printf("%v\n", useful_things.Eval(eq))
}
