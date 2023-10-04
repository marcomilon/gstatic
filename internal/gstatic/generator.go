package gstatic

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type generator struct {
	layoutPath string
	sourcePath string
	targetPath string
}

type task struct {
	generator  generator
	targetPath string
}

type copyFileWorker task
type layoutRenderWorker task
type simpleRenderWorker task

func (copyFileWorker copyFileWorker) execute() error {
	targetAssetname := getTargetDirname(copyFileWorker.generator.sourcePath, copyFileWorker.targetPath)
	targetFile := filepath.Join(copyFileWorker.generator.targetPath, targetAssetname)

	return copyFile(copyFileWorker.targetPath, targetFile)
}

func (layoutRenderWorker layoutRenderWorker) execute() error {
	targetAssetname := getTargetDirname(layoutRenderWorker.generator.sourcePath, layoutRenderWorker.targetPath)
	targetFile := filepath.Join(layoutRenderWorker.generator.targetPath, targetAssetname)
	tpl := gstaticLayoutTpl{layoutRenderWorker.generator.layoutPath, layoutRenderWorker.targetPath, targetFile}
	return tpl.render()
}

func (simpleRenderWorker simpleRenderWorker) execute() error {
	tpl := gstaticSimpleTpl{simpleRenderWorker.generator.sourcePath, simpleRenderWorker.generator.targetPath}
	return tpl.render()
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

	var wg sync.WaitGroup

	layoutFile, err := findLayout(srcFolder)
	if err != nil {
		return fmt.Errorf("[%s] %v", "generate", err)
	}

	generator := generator{layoutFile, srcFolder, targetFolder}

	resolver := resolver(generator, &wg)

	err = filepath.Walk(srcFolder, resolver)
	wg.Wait()

	return err

}

func resolver(generator generator, wg *sync.WaitGroup) filepath.WalkFunc {

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

				go func() {
					defer wg.Done()
					wg.Add(1)
					identifier := fmt.Sprintf("layoutRenderWorker: %s", path)
					executor(identifier, layoutRenderWorker{generator, path}, false)
				}()

				return nil

			}

			go func() {
				defer wg.Done()
				wg.Add(1)
				identifier := fmt.Sprintf("simpleRenderWorker: %s", path)
				executor(identifier, simpleRenderWorker{generator, path}, false)
			}()

			return nil

		}

		go func() {
			defer wg.Done()
			wg.Add(1)
			identifier := fmt.Sprintf("copyFileWorker: %s", path)
			executor(identifier, copyFileWorker{generator, path}, false)
		}()

		return nil

	}

}

func getTargetDirname(srcFolder, path string) string {
	s := strings.Replace(path, srcFolder, "", 1)
	return strings.TrimLeft(s, "/")
}
