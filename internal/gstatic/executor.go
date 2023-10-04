package gstatic

import (
	"fmt"
	"time"
)

type elapsedTime struct {
	identifier string
	start      time.Time
	end        time.Time
}

type worker interface {
	execute() error
}

func executor(identifier string, w worker, debug bool) {
	var err error

	if debug {
		enlapsedTime := StartTimer(identifier)
		err = w.execute()
		EndTimer(enlapsedTime)
	} else {
		err = w.execute()
	}

	if err != nil {
		fmt.Printf("Something went wrong: %v\n", err.Error())
	}
}

func StartTimer(identifier string) elapsedTime {
	return elapsedTime{identifier, time.Now(), time.Time{}}
}

func EndTimer(elapsedTime elapsedTime) {
	elapsedTime.end = time.Now()
	elapsed := elapsedTime.end.Sub(elapsedTime.start)
	fmt.Printf("[%s]: %v\n", elapsedTime.identifier, elapsed)
}
