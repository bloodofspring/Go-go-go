package kata

import "fmt"

func IpsBetween(start, end string) int {
	var startNum, endNum int
	var startOctets, endOctets [4]int

	fmt.Sscanf(start, "%d.%d.%d.%d", &startOctets[0], &startOctets[1], &startOctets[2], &startOctets[3])
	fmt.Sscanf(end, "%d.%d.%d.%d", &endOctets[0], &endOctets[1], &endOctets[2], &endOctets[3])

	for i := 0; i < 4; i++ {
		startNum = startNum<<8 + startOctets[i]
		endNum = endNum<<8 + endOctets[i]
	}

	return endNum - startNum
}
