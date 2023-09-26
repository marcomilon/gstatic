package gstatic

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var templateTestTargetFolder string = os.TempDir() + "gstatictest"

func TestSimpleTemplate(t *testing.T) {

	setupTemplateTest(t)

	sourcePath := "testdata/simpletpl/index.html"
	targetPath := filepath.Join(templateTestTargetFolder, "index.html")

	tpl := gstaticSimpleTpl{sourcePath, targetPath}
	err := tpl.render()
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	_, err = os.Stat(targetPath)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	result, err := os.ReadFile(targetPath)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}
	sectionResult := string(result)
	sectionExpected := "<p>Hello world</p>"
	if !strings.EqualFold(sectionResult, sectionExpected) {
		t.Errorf("expected %s; got %s", sectionExpected, sectionResult)
	}

}

func TestLayoutTemplate(t *testing.T) {

	setupTemplateTest(t)

	layoutPath := "testdata/layouttpl/layout/layout.html"
	sourcePath := "testdata/layouttpl/index.html"
	targetPath := filepath.Join(templateTestTargetFolder, "index.html")

	tpl := gstaticLayoutTpl{layoutPath, sourcePath, targetPath}
	err := tpl.render()
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	_, err = os.Stat(targetPath)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	result, err := os.ReadFile(targetPath)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	sectionResult := string(result)
	sectionExpected := "<h1>Hello world</h1><main><p>Index</p></main>"
	if !strings.EqualFold(sectionResult, sectionExpected) {
		t.Errorf("expected %v; got %v", sectionExpected, sectionResult)
	}

}

func setupTemplateTest(t *testing.T) {
	log.SetOutput(io.Discard)
	files, err := filepath.Glob(filepath.Join(templateTestTargetFolder, "*"))
	if err != nil {
		t.Fatal("Unable to setup test")
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			t.Fatal("Unable to setup test")
		}
	}

	os.Remove(templateTestTargetFolder)

	os.MkdirAll(templateTestTargetFolder, 0755)
}
