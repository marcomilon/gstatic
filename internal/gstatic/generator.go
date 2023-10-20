package gstatic

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
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

func copyFileWorker(t task) error {
	targetAssetname := getTargetDirname(t.generator.sourcePath, t.targetPath)
	targetFile := filepath.Join(t.generator.targetPath, targetAssetname)

	return copyFile(t.targetPath, targetFile)
}

func layoutRenderWorker(t task) error {
	targetAssetname := getTargetDirname(t.generator.sourcePath, t.targetPath)
	targetFile := filepath.Join(t.generator.targetPath, targetAssetname)
	tpl := gstaticLayoutTpl{t.generator.layoutPath, t.targetPath, targetFile}
	return tpl.render()
}

func simpleRenderWorker(t task) error {
	targetAssetname := getTargetDirname(t.generator.sourcePath, t.targetPath)
	targetFile := filepath.Join(t.generator.targetPath, targetAssetname)
	tpl := gstaticSimpleTpl{t.targetPath, targetFile}
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
	eg := errgroup.Group{}

	layoutFile, err := findLayout(srcFolder)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	generator := generator{layoutFile, srcFolder, targetFolder}

	resolver := resolver(generator, &wg, &eg)

	err = filepath.Walk(srcFolder, resolver)
	wg.Wait()

	errGroup := eg.Wait()
	if errGroup != nil {
		return errGroup
	}

	return err

}

func resolver(generator generator, wg *sync.WaitGroup, eg *errgroup.Group) filepath.WalkFunc {

	useLayout := generator.layoutPath != ""

	return func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return fmt.Errorf("%v", err)
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

		task := task{generator, path}

		if ext == ".html" {

			if isTplPlainHtml(path) {

				wg.Add(1)
				go func() {
					defer wg.Done()
					copyFileWorker(task)
				}()

				return nil

			}

			if useLayout {
				wg.Add(1)
				go func() {
					defer wg.Done()
					layoutRenderWorker(task)
				}()

				return nil

			}

			eg.Go(func() error {
				return simpleRenderWorker(task)
			})

			return nil

		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			copyFileWorker(task)
		}()

		return nil

	}

}

func isTplPlainHtml(sourceFile string) bool {
	s := strings.Replace(sourceFile, ".html", "", 1)
	sourceFile = s + ".yaml"
	_, err := os.Stat(sourceFile)
	return os.IsNotExist(err)

}

func getTargetDirname(srcFolder, path string) string {
	s := strings.Replace(path, srcFolder, "", 1)
	return strings.TrimLeft(s, "/")
}
