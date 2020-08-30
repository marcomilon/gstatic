package generator

import "io"

// VarReader is an interface which goal is to extract variable from a data source. For example a Yaml file, Json file or even DB
type VarReader interface {
	GetVarsForTpl(r io.Reader) (map[interface{}]interface{}, error)
	GetDsExtension() string
}
