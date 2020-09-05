package main

import (
	"fmt"
	"os"

	"github.com/marcomilon/gstatic/internal/datasource"
	"github.com/marcomilon/gstatic/internal/gstatic"
)

func main() {

	if len(os.Args) < 3 {
		usage()
	}

	argsWithoutProg := os.Args[1:]

	srcFolder := argsWithoutProg[0]
	targetFolder := argsWithoutProg[1]

	ds := datasource.Yaml{}

	yamlGen := gstatic.Generator{VarReader: ds}

	err := yamlGen.Generate(srcFolder, targetFolder)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Something went wrong: %v\n", err.Error())
	}

}

func usage() {
	fmt.Println("Usage: gstatic <sourceFolder> <targetFolder>")
	os.Exit(0)
}
