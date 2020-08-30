package datasource_test

import (
	"strings"
	"testing"

	"github.com/marcomilon/gstatic/internal/datasource"
)

func TestGetVarsForTpl(t *testing.T) {

	s := "hello: Hello world"

	r := strings.NewReader(s)

	yamlExtractor := datasource.Yaml{}

	m, err := yamlExtractor.GetVarsForTpl(r)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	if m["hello"] != "Hello world" {
		t.Errorf("expected %v; got %v", "Hello world", m["hello"])
	}

}

func TestGetDsExtension(t *testing.T) {
	yamlExtractor := datasource.Yaml{}

	e := yamlExtractor.GetDsExtension()
	if e != ".yaml" {
		t.Errorf("expected %v; got %v", "yaml", e)
	}
}
