package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/marcomilon/gstatic/internal/datasource"
	"github.com/marcomilon/gstatic/internal/gstatic"
)

func main() {

	layoutPtr := flag.String("layout", "layout/layout.html", "Path to the layout")
	basePtr := flag.String("base", "base", "Name of the main template")

	flag.Parse()

	config := gstatic.Config{
		Layout: *layoutPtr,
		Base:   *basePtr,
	}

	if len(os.Args) < 3 {
		usage()
	}

	argsWithoutProg := os.Args[1:]

	srcFolder := argsWithoutProg[0]
	targetFolder := argsWithoutProg[1]

	ds := datasource.Yaml{}

	yamlGen := gstatic.Generator{
		Config:    config,
		VarReader: ds,
	}

	err := yamlGen.Generate(srcFolder, targetFolder)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Something went wrong: %v\n", err.Error())
		os.Exit(1)
	}

}

func usage() {

	fmt.Println("Usage: gstatic <sourceFolder> <targetFolder>")
	fmt.Println("Use -h for help")
	os.Exit(1)

}
