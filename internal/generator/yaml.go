package generator

import (
	"io"
	"os"
	"path/filepath"
	"strings"
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
			if err := copyFile(path, target); err != nil {
				return err
			}
		} else {
			if err := createDir(target); err != nil {
				return err
			}
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

	// Create new file
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
