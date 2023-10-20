package gstatic

import (
	"io"

	"gopkg.in/yaml.v2"
)

func getVarsForTpl(r io.Reader) (map[interface{}]interface{}, error) {

	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	return m, nil

}
