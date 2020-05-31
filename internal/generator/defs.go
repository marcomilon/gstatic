package generator

const layoutFolder = "/layout"
const templateExt = ".html"
const sourceExt = ".yaml"
const baseTemplate = "base"

// VarReader is an interface which goal is to extract variable from a data source. For example a Yaml file, Json file or even DB
type VarReader interface {
	GetVarsForTpl(tpl string) (map[interface{}]interface{}, error)
}
