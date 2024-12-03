package useful_things

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func RandomDate() string {
	var startYr, startMnthInt, startDay, tab int

	fmt.Print("Enter start date [year/month/day]: ")
	_, err := fmt.Scanf("%d/%d/%d", &startYr, &startMnthInt, &startDay)
	fmt.Print("Enter tab before last gen date and now [days]: ")
	_, err1 := fmt.Scanf("%d", &tab)

	if err != nil || err1 != nil {
		return "<Incorrect input.>"
	}

	now := time.Now()
	minTime := time.Date(startYr, time.Month(startMnthInt), startDay, now.Hour(), now.Minute(), now.Second(), 0, now.Location()).Unix()
	maxTime := now.Unix() - int64(tab*24*60*60)
	delta := maxTime - minTime
	sec := rand.Int64N(delta) + minTime

	return time.Unix(sec, 0).String()
}
