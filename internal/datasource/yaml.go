package datasource

import (
	"path/filepath"
	"strings"

	"github.com/marcomilon/gstatic/internal/helpers"
	"gopkg.in/yaml.v2"
)

// Yaml a struct that get variables from a Yaml file
type Yaml struct{}

// GetVarsForTpl is a implementation for VarReader. This implementatin will get variables from a Yaml file
func (Yaml) GetVarsForTpl(tpl string) (map[interface{}]interface{}, error) {

	sourcePath := strings.TrimSuffix(tpl, filepath.Ext(tpl))
	dataSource := sourcePath + ".yaml"

	data, err := helpers.ReadFile(dataSource)
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
