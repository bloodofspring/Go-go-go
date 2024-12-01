package useful_things

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunEval() {
	var eq string

	fmt.Println("input equation:")
	reader := bufio.NewReader(os.Stdin)
	eq, err := reader.ReadString('\n')

	if err != nil {
		return
	}

	fmt.Printf("Result: %v\n", Eval(eq))
}

func Eval(equation string) int {
	fmt.Printf("Evaluating equation: %s\n", equation)

	equation = strings.Trim(equation, "\n") // Like strip in python
	data := strings.Split(equation, " ")

	firstNum, _ := strconv.Atoi(data[0])
	operation := data[1]
	SecondNum, _ := strconv.Atoi(data[2])

	switch operation {
	case "+":
		return firstNum + SecondNum
	case "-":
		return firstNum - SecondNum
	case "*":
		return firstNum * SecondNum
	case "/":
		return firstNum / SecondNum
	case "%":
		return firstNum % SecondNum
	default:
		return -1
	}
}
