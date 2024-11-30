package archive

import "fmt"

func CycleTestFunc() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("I value: %d\n", i)
	}

	for {
		println("Бесконечный цикл")
		// break (also can pass continue)
	}
}
