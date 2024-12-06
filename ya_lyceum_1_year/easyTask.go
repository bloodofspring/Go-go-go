package ya_lyceum_1_year

import "fmt"

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func IqTest() {
	var nPeople, iq, iqSum int
	var err error

	iqSum = 0

	fmt.Print("Enter number of people: ")
	_, err = fmt.Scanf("%d", &nPeople)
	handleError(err)

	for i := 0; i < nPeople; i++ {
		var avgIq int
		if i != 0 {
			avgIq = iqSum / i
		} else {
			avgIq = 0
		}

		fmt.Printf("[%d] Enter user's IQ: ", i)
		_, err = fmt.Scanf("%d", &iq)
		handleError(err)

		iqSum += iq

		if i == 0 {
			fmt.Println("0")
		} else if iq < avgIq {
			fmt.Println("<")
		} else if iq > avgIq {
			fmt.Println(">")
		} else {
			fmt.Println("0")
		}

	}

}
