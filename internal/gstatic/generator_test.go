package gstatic

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var generatorTestTargetFolder string = os.TempDir() + "gstatictest"

func TestGenerateSimple(t *testing.T) {

	generatorTemplateTest(t)

	sourcePath := "testdata/simpletpl"
	targetPath := generatorTestTargetFolder

	err := Generate(sourcePath, targetPath)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	targetFile := filepath.Join(generatorTestTargetFolder, "about.html")
	_, err = os.Stat(targetFile)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

}

func TestGenerateLayout(t *testing.T) {

	generatorTemplateTest(t)

	sourcePath := "testdata/layouttpl"
	targetPath := generatorTestTargetFolder

	err := Generate(sourcePath, targetPath)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

}

func TestIsTplPlainHtml(t *testing.T) {

	sourceFile := "testdata/simpletpl/index.html"
	hasSource := isTplPlainHtml(sourceFile)
	if hasSource {
		t.Errorf("expected %v; got %v", false, hasSource)
	}

	sourceFile2 := "testdata/simpletpl/about.html"
	hasSource2 := isTplPlainHtml(sourceFile2)
	if !hasSource2 {
		t.Errorf("expected %v; got %v", true, hasSource)
	}

}

func TestFindLayout(t *testing.T) {

	var source string
	var layoutFile string
	var err error

	source = "testdata/layouttpl"

	layoutFile, err = findLayout(source)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	if layoutFile != "testdata/layouttpl/layout/layout.html" {
		t.Errorf("expected %v; got %v", true, layoutFile)
	}

	source = "testdata/simpletpl"

	layoutFile, err = findLayout(source)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	if layoutFile != "" {
		t.Errorf("expected %v; got %v", "", layoutFile)
	}

}

func generatorTemplateTest(t *testing.T) {
	log.SetOutput(io.Discard)
	files, err := filepath.Glob(filepath.Join(generatorTestTargetFolder, "*"))
	if err != nil {
		t.Fatal("Unable to setup test")
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			t.Fatal("Unable to setup test")
		}
	}

	os.Remove(generatorTestTargetFolder)

	os.MkdirAll(generatorTestTargetFolder, 0755)
}
