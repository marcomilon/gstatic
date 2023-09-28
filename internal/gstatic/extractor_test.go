package gstatic

import (
	"strings"
	"testing"
)

func TestGetVarsForTpl(t *testing.T) {

	s := "hello: Hello world"

	r := strings.NewReader(s)

	m, err := getVarsForTpl(r)
	if err != nil {
		t.Errorf("expected %v; got %v", nil, err)
	}

	if m["hello"] != "Hello world" {
		t.Errorf("expected %v; got %v", "Hello world", m["hello"])
	}

}
