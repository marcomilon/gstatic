package generator

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

const layoutFolder = "/layout"
const templateExt = ".html"
const sourceExt = ".yaml"
const baseTemplate = "base"

// Yaml is a resolver that will use a yaml file to hold the variables
type Yaml struct {
	Source string
	Target string
}

func (yaml Yaml) resolver() filepath.WalkFunc {

	layouts, err := readLayouts(yaml.source() + layoutFolder)
	if err != nil {
		log.Printf("Warning: layouts not found, %v", err)
	}

	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		targetFile := strings.Replace(path, yaml.source(), "", 1)
		target := yaml.target() + targetFile

		if !info.IsDir() {

			template := strings.TrimSuffix(path, filepath.Ext(path))
			source := template + sourceExt

			if _, err := os.Stat(source); os.IsNotExist(err) {
				return copyFile(path, target)
			}

			parseTpl := filepath.Ext(path) == templateExt && filepath.Dir(path) != yaml.source()+layoutFolder
			if parseTpl {
				if err := parseFile(layouts, source, target, path); err != nil {
					return err
				}
			}

			return nil
		}

		if err := createDir(target); err != nil {
			return err
		}

		return nil
	}
}

func (yaml Yaml) source() string {
	return yaml.Source
}

func (yaml Yaml) target() string {
	return yaml.Target
}

func createDir(target string) error {
	err := os.MkdirAll(target, 0755)
	if err != nil {
		return err
	}

	return nil
}

func copyFile(source string, target string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	newFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer newFile.Close()

	if _, err := io.Copy(newFile, sourceFile); err != nil {
		return err
	}

	return nil
}

func writeFile(source []byte, target string) error {

	err := ioutil.WriteFile(target, source, 0644)
	if err != nil {
		return err
	}

	return nil
}

func readLayouts(path string) ([]string, error) {
	var files []string

	filesInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range filesInfo {
		ext := filepath.Ext(file.Name())
		if ext == templateExt {
			files = append(files, path+"/"+file.Name())
		}
	}

	return files, nil

}

func parseFile(layouts []string, source string, target string, tpl string) error {

	var html []string

	data, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	if len(layouts) > 0 {
		html = append(layouts, tpl)
	} else {
		html = append(html, tpl)
	}

	tmpl, err := template.ParseFiles(html...)
	if err != nil {
		return err
	}

	if len(layouts) > 0 {
		err = tmpl.ExecuteTemplate(f, baseTemplate, m)
	} else {
		err = tmpl.Execute(f, m)
	}
	if err != nil {
		return err
	}

	return nil
}
