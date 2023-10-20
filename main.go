package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/marcomilon/gstatic/internal/gstatic"
)

type elapsedTime struct {
	identifier string
	start      time.Time
	end        time.Time
}

func main() {

	if len(os.Args) < 3 {
		usage()
	}
	argsWithoutProg := os.Args[1:]

	srcFolder := argsWithoutProg[0]
	targetFolder := argsWithoutProg[1]

	elapsedTime := startTimer("gstatic")
	err := gstatic.Generate(srcFolder, targetFolder)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Something went wrong: %v\n", err.Error())
		os.Exit(1)
	}
	endTimer(elapsedTime)

}

func startTimer(identifier string) elapsedTime {
	return elapsedTime{identifier, time.Now(), time.Time{}}
}

func endTimer(elapsedTime elapsedTime) {
	elapsedTime.end = time.Now()
	elapsed := elapsedTime.end.Sub(elapsedTime.start)
	fmt.Printf("[%s]: %v\n", elapsedTime.identifier, elapsed)
}

func usage() {

	fmt.Println("Usage: gstatic <sourceFolder> <targetFolder>")
	flag.PrintDefaults()
	os.Exit(1)

}
