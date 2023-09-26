package gstatic

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var fileTestTargetFolder string = os.TempDir() + "gstatictest"

func TestCopyFile(t *testing.T) {

	setupFileTest(t)

	source := "testdata/simpletpl/index.html"
	target := filepath.Join(fileTestTargetFolder, "index.html")

	err := copyFile(source, target)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

}

func TestCreateTargetFolder(t *testing.T) {

	setupFileTest(t)

	target := filepath.Join(fileTestTargetFolder, "newfolder")

	err := createTargetFolder(target)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

}

func setupFileTest(t *testing.T) {
	log.SetOutput(io.Discard)
	files, err := filepath.Glob(filepath.Join(fileTestTargetFolder, "*"))
	if err != nil {
		t.Fatal("Unable to setup test")
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			t.Fatal("Unable to setup test")
		}
	}

	os.Remove(fileTestTargetFolder)

	os.MkdirAll(fileTestTargetFolder, 0755)
}
