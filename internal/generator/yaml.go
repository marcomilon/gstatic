package generator

import (
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// Yaml is a resolver that will use a yaml file to hold the variables
type Yaml struct {
	Source string
	Target string
}

func (yaml Yaml) resolver() filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err

		}

		targetFile := strings.Replace(path, yaml.source(), "", 1)
		target := yaml.target() + targetFile

		if !info.IsDir() {

			template := strings.TrimSuffix(path, filepath.Ext(path))
			source := template + ".yaml"

			if _, err := os.Stat(source); os.IsNotExist(err) {
				return copyFile(path, target)
			}

			err := parseFile(source, target, path)
			if err != nil {
				return err
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

func parseFile(source string, target string, tpl string) error {

	data, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(tpl)
	if err != nil {
		return err
	}

	f, err := os.Create(target)
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, m)
	if err != nil {
		return err
	}

	f.Close()

	return nil
}
