package datasource

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Yaml a struct that get variables from a Yaml file
type Yaml struct{}

// GetVarsForTpl is a implementation for VarReader. This implementatin will get variables from a Yaml file
func (Yaml) GetVarsForTpl(r io.Reader) (map[interface{}]interface{}, error) {

	data, err := ioutil.ReadAll(r)
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

// GetDsExtension returns the file extension of the data source
func (Yaml) GetDsExtension() string {
	return ".yaml"
}
