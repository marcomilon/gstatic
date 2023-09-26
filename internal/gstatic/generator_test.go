package gstatic

import "testing"

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
