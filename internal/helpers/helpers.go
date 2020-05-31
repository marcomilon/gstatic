package helpers

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const templateExt = ".html"

func CreateDir(target string) error {
	err := os.MkdirAll(target, 0755)
	if err != nil {
		return err
	}

	return nil
}

func CopyFile(source string, target string) error {
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

func WriteFile(source []byte, target string) error {

	err := ioutil.WriteFile(target, source, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile(source string) ([]byte, error) {
	data, err := ioutil.ReadFile(source)
	if err != nil {
		return nil, err
	}

	return data, err
}

func ReadLayouts(path string) ([]string, error) {
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

//func getDataSourcePath()
