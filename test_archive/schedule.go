package test_archive

// to install package: go get <name>@version
// to uninstall package: go get <name>@none
// type Args[E any] []E
// import "as": import yourName "way"

import (
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"time"
)

func job(someInteger int, someS string) int {
	fmt.Println(someS)
	return someInteger + 1
}

func ScheduleSample() {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	j, err := scheduler.NewJob(gocron.DurationJob(10*time.Second), gocron.NewTask(job, 1, "Hello World!"))
	if err != nil {
		panic(err)
	}

	fmt.Println(j.ID())
	scheduler.Start()

	select { // wait a minute
	case <-time.After(time.Minute):
	}

	err = scheduler.Shutdown()
	if err != nil {
		panic(err)
	}
}

func SchedularMain() {
	ScheduleSample()
}
