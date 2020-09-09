package gstatic

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var dirnametests = []struct {
	srcFolder string
	path      string
	expected  string
}{
	{"testdata/basic", "testdata/basic", ""},
	{"testdata/basic", "testdata/basic/section", "section"},
	{"testdata/basic", "testdata/basic/section/subsection", "section/subsection"},
}

var targetFolder string = os.TempDir() + "gstatictest"

func TestGetTargetDirname(t *testing.T) {

	for _, tt := range dirnametests {
		t.Run(tt.path, func(t *testing.T) {
			got := getTargetDirname(tt.srcFolder, tt.path)
			if tt.expected != got {
				t.Errorf("expected %v; got %v", tt.expected, got)
			}
		})
	}

}

func TestCopyAsset(t *testing.T) {

	setup(t)

	source := "testdata/basic/350x150.png"
	target := targetFolder + "/350x150.png"

	err := copyAsset(source, target)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	if _, err := os.Stat(target); os.IsNotExist(err) {
		t.Errorf("expected %v; got %v", nil, err)
	}

}

func TestGetSourceFilename(t *testing.T) {

	path := "testdata/basic/index.html"
	expected := "testdata/basic/index.yaml"

	sourceFilename := getSourceFilename(path, ".yaml")
	if sourceFilename != expected {
		t.Errorf("expected %v; got %v", expected, sourceFilename)
	}

}

func TestMergeSourceFiles(t *testing.T) {

	source1 := "testdata/source/source1.yaml"
	source2 := "testdata/source/source2.yaml"
	expected := `title: hello
body: world
footer: goodbye`

	mergeSourceReader, err := mergeSourceFile(source1, source2)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	mergeSourceBytes, err2 := ioutil.ReadAll(mergeSourceReader)
	if err2 != nil {
		t.Errorf("expected %v; got %v", nil, err2)
	}

	mergeSource := string(mergeSourceBytes)

	if mergeSource != expected {
		t.Errorf("expected %v; got %v", expected, mergeSource)
	}

}

func TestHasSourceFilename(t *testing.T) {
	path := "testdata/static/index.html"

	hasSourceFile := hasSourceFilename(path, ".yaml")
	if hasSourceFile {
		t.Errorf("expected %v; got %v", false, hasSourceFile)
	}

	path2 := "testdata/basic/index.html"
	hasSourceFile2 := hasSourceFilename(path2, ".yaml")
	if !hasSourceFile2 {
		t.Errorf("expected %v; got %v", false, hasSourceFile2)
	}
}

func TestValidateTargetFolder(t *testing.T) {

	path := "testdata/basic/"
	err := validateTargetFolder(path)
	if err == nil {
		t.Errorf("expected %v; got %v", err, nil)
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
