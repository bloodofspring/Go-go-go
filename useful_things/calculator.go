package useful_things

import (
	"fmt"
	"strconv"
	"strings"
)

func Eval(equation string) int {
	fmt.Printf("Evaluating equation: %s\n", equation)
	data := [3]string{}

	equation = strings.Trim(equation, "\n")
	copy(data[:], strings.Split(equation, " "))

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
		return 1
	}
}