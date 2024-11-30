package archive

import (
	"fmt"
	"strings"
)

func MissivesTest() { // <name>(<arg1> <type>, <arg2> <type>) (<return_value> <type>) {}
	var someArray = [4]int{7, 7, 7, 7}
	fmt.Printf("%v", someArray)

	var UnknownLengthArray = [...]int{9, 9, 9, 9}
	fmt.Printf("%v", UnknownLengthArray)

	var EmptyArray = [...]int{}
	fmt.Printf("%v", EmptyArray)

	abcUpper := strings.ToUpper("abs")
	fmt.Printf("%v", abcUpper)

	return // return statement
}
