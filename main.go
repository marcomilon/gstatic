package main

import (
	"fmt"
	"os"

	"github.com/marcomilon/gstatic/internal/datasource"
	"github.com/marcomilon/gstatic/internal/generator"
)

func main() {
	fmt.Println("gStatic")

	argsWithoutProg := os.Args[1:]

	srcFolder := argsWithoutProg[0]
	targetFolder := argsWithoutProg[1]

	ds := datasource.Yaml{}

	yamlGen := generator.Generator{ds}

	err := yamlGen.Generate(srcFolder, targetFolder)
	if err != nil {
		fmt.Printf("expected %v; got %v", nil, err)
	}

}
