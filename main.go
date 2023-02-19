package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/marcomilon/gstatic/internal/datasource"
	"github.com/marcomilon/gstatic/internal/gstatic"
)

func main() {

	layoutPtr := flag.String("layout", "layout/layout.html", "Path to the layout in the target folder")
	basePtr := flag.String("base", "base", "Name of the main template")
	forcePtr := flag.Bool("force", false, "Remove all file on targetfolder and run website generator")

	if len(os.Args) < 3 {
		usage()
	}

	flag.CommandLine.Parse(os.Args[3:])

	config := gstatic.Config{
		Layout:                 *layoutPtr,
		Base:                   *basePtr,
		ForceWebSiteGeneration: *forcePtr,
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
	flag.PrintDefaults()
	os.Exit(1)

}
