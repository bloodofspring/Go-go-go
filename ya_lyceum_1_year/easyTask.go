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

// Another task

func SchwarzeneggerVsGodzilla() {
	var shots int
	var damage float32
	var err error

	fmt.Print("Enter number of shots: ")
	_, err = fmt.Scanf("%d", &shots)
	handleError(err)

	for i := 0; i < shots; i++ {
		var numerator, denominator int
		fmt.Printf("[%d] Enter damage (a/b): ", i)
		_, err = fmt.Scanf("%d/%d", &numerator, &denominator)
		handleError(err)

		damage += float32(numerator) / float32(denominator)
	}

	fmt.Printf("%.2f", damage)
}
