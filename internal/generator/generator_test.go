package generator_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/marcomilon/gstatic/internal/generator"
)

var out string = os.TempDir() + "gstatictest/public"

func TestGenerator(t *testing.T) {

	setup(t)

	yaml := generator.Yaml{"testdata/yaml", out}
	err := generator.Generate(yaml)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	index := out + string(os.PathSeparator) + "index.html"
	if _, err := os.Stat(index); err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	section := out + string(os.PathSeparator) + "/section/section.html"
	if _, err := os.Stat(section); err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}
}

func setup(t *testing.T) {
	files, err := filepath.Glob(filepath.Join(out, "*"))
	if err != nil {
		t.Fatal("Unable to setup test")
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			t.Fatal("Unable to setup test")
		}
	}

	os.Remove(out)
}
