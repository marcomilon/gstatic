package gstatic_test

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/marcomilon/gstatic/internal/datasource"
	"github.com/marcomilon/gstatic/internal/gstatic"
)

var targetFolder string = os.TempDir() + "gstatictest"
var srcFolderBasic = "testdata/basic"
var srcFolderComposition = "testdata/composition"
var srcFolderStatic = "testdata/static"

func TestBasicGenerator(t *testing.T) {

	setup(t)

	ds := datasource.Yaml{}

	config := gstatic.Config{
		"layout/layout.html",
		"base",
	}

	yamlGen := gstatic.Generator{
		config,
		ds,
	}

	err := yamlGen.Generate(srcFolderBasic, targetFolder)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	index := targetFolder + string(os.PathSeparator) + "index.html"
	indexTpl, err := ioutil.ReadFile(index)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	indexResult := string(indexTpl)
	indexExpected := "<p>Hello world</p>"
	if strings.ToLower(indexResult) != strings.ToLower(indexExpected) {
		t.Errorf("expected %v; got %v", indexExpected, indexResult)
	}

	section := targetFolder + string(os.PathSeparator) + "section/section.html"
	sectionTpl, err := ioutil.ReadFile(section)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	sectionResult := string(sectionTpl)
	sectionExpected := "<p>Marco</p>"
	if strings.ToLower(sectionResult) != strings.ToLower(sectionExpected) {
		t.Errorf("expected %v; got %v", sectionExpected, sectionResult)
	}
}

func TestCompositionGenerator(t *testing.T) {

	setup(t)

	ds := datasource.Yaml{}

	config := gstatic.Config{
		"layout/layout.html",
		"base",
	}

	yamlGen := gstatic.Generator{
		config,
		ds,
	}

	err := yamlGen.Generate(srcFolderComposition, targetFolder)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	index := targetFolder + string(os.PathSeparator) + "index.html"
	indexTpl, err := ioutil.ReadFile(index)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	indexResult := strings.TrimSpace(string(indexTpl))
	indexExpected := "<h1>Hello world</h1><main><p>Index</p></main>"
	if strings.ToLower(indexResult) != strings.ToLower(indexExpected) {
		t.Errorf("expected %v; got %v", indexExpected, indexResult)
	}

}

func TestStaticGenerator(t *testing.T) {

	setup(t)

	ds := datasource.Yaml{}

	config := gstatic.Config{
		"layout/layout.html",
		"base",
	}

	yamlGen := gstatic.Generator{
		config,
		ds,
	}

	err := yamlGen.Generate(srcFolderStatic, targetFolder)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	index := targetFolder + string(os.PathSeparator) + "index.html"
	indexTpl, err := ioutil.ReadFile(index)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	indexResult := strings.TrimSpace(string(indexTpl))
	indexExpected := "<p>static</p>"
	if strings.ToLower(indexResult) != strings.ToLower(indexExpected) {
		t.Errorf("expected %v; got %v", indexExpected, indexResult)
	}

}

func setup(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	files, err := filepath.Glob(filepath.Join(targetFolder, "*"))
	if err != nil {
		t.Fatal("Unable to setup test")
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			t.Fatal("Unable to setup test")
		}
	}

	os.Remove(targetFolder)
	os.MkdirAll(targetFolder, 0755)
}
