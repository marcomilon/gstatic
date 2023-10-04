package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/marcomilon/gstatic/internal/gstatic"
)

func main() {

	if len(os.Args) < 3 {
		usage()
	}
	argsWithoutProg := os.Args[1:]

	srcFolder := argsWithoutProg[0]
	targetFolder := argsWithoutProg[1]

	elapsedTime := gstatic.StartTimer("gstatic")
	err := gstatic.Generate(srcFolder, targetFolder)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Something went wrong: %v\n", err.Error())
		os.Exit(1)
	}
	gstatic.EndTimer(elapsedTime)
}

func usage() {

	fmt.Println("Usage: gstatic <sourceFolder> <targetFolder>")
	flag.PrintDefaults()
	os.Exit(1)

}
