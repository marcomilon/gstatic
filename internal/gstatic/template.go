package gstatic

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/marcomilon/gstatic/internal/datasource"
)

type gstaticSimpleTpl struct {
	sourcePath string
	targetPath string
}

type gstaticLayoutTpl struct {
	layoutPath string
	sourcePath string
	targetPath string
}

func (tpl gstaticLayoutTpl) render() error {

	dataSource := getSourceFilename(tpl.sourcePath)
	layoutSource := getSourceFilename(tpl.layoutPath)

	variablesFile, err := mergeFiles(dataSource, layoutSource)
	if err != nil {
		return fmt.Errorf("[%s] %v", "render", err)
	}

	variables, err := extractVariables(variablesFile)
	if err != nil {
		log.Printf("%s: %v\n", "gstaticLayoutTpl unable to extract variables", err)
		return fmt.Errorf("[%s] %v", "render", err)
	}

	return parse([]string{tpl.layoutPath, tpl.sourcePath}, tpl.targetPath, variables)
}

func (tpl gstaticSimpleTpl) render() error {

	dataSource := getSourceFilename(tpl.sourcePath)

	dataSourceFile, err := os.Open(dataSource)
	if err != nil {
		return fmt.Errorf("[%s] %v", "render", err)
	}
	defer dataSourceFile.Close()

	variables, err := extractVariables(dataSourceFile)
	if err != nil {
		return fmt.Errorf("[%s] %v", "render", err)
	}

	return parse([]string{tpl.sourcePath}, tpl.targetPath, variables)

}

func parse(sourcePath []string, targetPath string, variables map[interface{}]interface{}) error {

	targetFile, err := os.Create(targetPath)
	if err != nil {
		return fmt.Errorf("[%s] %v", "parse", err)
	}
	defer targetFile.Close()

	tmpl, err := template.ParseFiles(sourcePath...)
	if err != nil {
		return fmt.Errorf("[%s] %v", "parse", err)
	}

	err = tmpl.Execute(targetFile, variables)
	if err != nil {
		return fmt.Errorf("[%s] %v", "parse", err)
	}

	return nil

}

func extractVariables(source io.Reader) (map[interface{}]interface{}, error) {
	ds := datasource.Yaml{}
	m, err := ds.GetVarsForTpl(source)
	if err != nil {
		return nil, fmt.Errorf("[%s] %v", "extractVariables", err)
	}

	return m, nil
}

func getSourceFilename(path string) string {
	filename := filepath.Base(path)
	extension := filepath.Ext(filename)
	sourceFilename := filename[0 : len(filename)-len(extension)]

	dirname := filepath.Dir(path)
	return dirname + string(os.PathSeparator) + sourceFilename + ".yaml"
}

func mergeFiles(path1, path2 string) (io.Reader, error) {

	content1, err := os.ReadFile(path1)
	if err != nil {
		return nil, fmt.Errorf("[%s] %v", "mergeFiles", err)
	}

	content2, err := os.ReadFile(path2)
	if err != nil {
		return nil, fmt.Errorf("[%s] %v", "mergeFiles", err)
	}

	mergeContent := string(content1) + "\n" + string(content2)

	return strings.NewReader(mergeContent), nil

}
