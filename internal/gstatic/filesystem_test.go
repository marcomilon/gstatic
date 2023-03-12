package gstatic

import (
	"fmt"
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

var publicFolder string = targetFolder + "/testdata/public"

func TestCleanTargetFolder(t *testing.T) {

	setup(t)
	dir := createTempDirWithFiles()

	err := cleanTargetFolder(dir)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

}

func TestValidateTargetFolder(t *testing.T) {

	setup(t)
	dir := createTempDirWithFiles()
	err := validateTargetFolder(dir)
	if err == nil {
		t.Errorf("expected %v; got %v", err, nil)
	}

	path := "testdata/404/"
	err2 := validateTargetFolder(path)
	if err2 == nil {
		t.Errorf("expected %v; got %v", err2, nil)
	}

}

func TestIsFolderEmpty(t *testing.T) {

	setup(t)

	dir := createTempEmptyDir(publicFolder)

	isEmpty, err := isFolderEmpty(dir)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	if isEmpty == false {
		t.Errorf("expected %v; got %v", true, err)
	}

	dir2 := createTempDirWithFiles()

	isEmpty2, err2 := isFolderEmpty(dir2)
	if err2 != nil {
		t.Errorf("expected %v; got %v", nil, err2)
	}

	if isEmpty2 == true {
		t.Errorf("expected %v; got %v", false, err)
	}

}

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

func createTempDirWithFiles() string {
	dir := createTempEmptyDir(publicFolder)

	filename := "example.txt"
	filepath := fmt.Sprintf("%s/%s", dir, filename)

	file, err := os.Create(filepath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	return dir
}

func createTempEmptyDir(dir string) string {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		panic(err.Error())
	}

	return dir
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
