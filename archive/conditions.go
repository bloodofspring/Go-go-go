package archive

func ConditionTest() {
	var firstNum, secondNum int
	firstNum = 123
	secondNum = firstNum + 23

	if firstNum < secondNum {
		println(firstNum)
	} else {
		println(secondNum)
	}

	var thirdNum int
	thirdNum = 900 + firstNum

	if firstNum < secondNum {
		println(firstNum)
	} else if thirdNum > secondNum {
		println(thirdNum)
	} else {
		println(secondNum)
	}

	switch firstNum {
	case secondNum:
		println(secondNum)
	case thirdNum:
		println(thirdNum)
	default:
		println("Nothing!")
	}
}
