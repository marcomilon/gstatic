package gstatic

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type generator struct {
	layoutPath string
	sourcePath string
	targetPath string
}

func findLayout(sourcePath string) (string, error) {

	sourceFile := filepath.Join(sourcePath, "/layout/layout.html")

	_, err := os.Stat(sourceFile)

	if err == nil {
		return sourceFile, nil
	} else if os.IsNotExist(err) {
		return "", nil
	} else {
		return "", err
	}

}

func Generate(srcFolder string, targetFolder string) error {

	layoutFile, err := findLayout(srcFolder)
	if err != nil {
		return fmt.Errorf("[%s] %v", "generate", err)
	}

	generator := generator{layoutFile, srcFolder, targetFolder}

	resolver := resolver(generator)

	return filepath.Walk(srcFolder, resolver)
}

func resolver(generator generator) filepath.WalkFunc {

	useLayout := generator.layoutPath != ""

	return func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return fmt.Errorf("[%s] %v", "resolver", err)
		}

		if path == generator.layoutPath {
			return nil
		}

		if info.IsDir() {

			targetFolder := getTargetDirname(generator.sourcePath, path)

			if targetFolder == "layout" {
				return nil
			}

			if targetFolder == "" {
				return nil
			}

			distFolder := filepath.Join(generator.targetPath, targetFolder)
			return createTargetFolder(distFolder)
		}

		ext := filepath.Ext(path)

		if ext == ".yaml" {
			return nil
		}

		if ext == ".html" {

			if useLayout {

				targetAssetname := getTargetDirname(generator.sourcePath, path)
				targetFile := filepath.Join(generator.targetPath, targetAssetname)
				tpl := gstaticLayoutTpl{generator.layoutPath, path, targetFile}
				return tpl.render()

			}

			tpl := gstaticSimpleTpl{generator.sourcePath, generator.targetPath}
			return tpl.render()

		}

		targetAssetname := getTargetDirname(generator.sourcePath, path)
		targetFile := filepath.Join(generator.targetPath, targetAssetname)

		return copyFile(path, targetFile)

	}

}

func getTargetDirname(srcFolder, path string) string {
	s := strings.Replace(path, srcFolder, "", 1)
	return strings.TrimLeft(s, "/")
}
