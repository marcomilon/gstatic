package gstatic

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
)

// Generator is the main struct. It use a VarReader to extract variables from a data source
type Generator struct {
	Config    Config
	VarReader VarReader
}

var (
	tplExtension = ".html"
)

// Generate is the main method used to generate a site
func (g Generator) Generate(srcFolder string, targetFolder string) error {

	if g.Config.ForceWebSiteGeneration {
		err := cleanTargetFolder(targetFolder)
		if err != nil {
			return err
		}

	}

	err := validateTargetFolder(targetFolder)
	if err != nil {
		return err
	}

	resolver := g.resolver(srcFolder, targetFolder)

	return filepath.Walk(srcFolder, resolver)
}

func (g Generator) resolver(srcFolder string, targetFolder string) filepath.WalkFunc {

	var useLayout bool = false
	var layout string = srcFolder + string(os.PathSeparator) + g.Config.Layout

	if _, err := os.Stat(layout); err == nil {
		useLayout = true
	}

	return func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return mkdir(srcFolder, targetFolder, path)
		}

		ext := filepath.Ext(path)
		if ext == tplExtension {
			targetFilename := getTargetDirname(srcFolder, path)

			sourceFile := hasSourceFilename(path, g.VarReader.GetDsExtension())
			if !sourceFile {
				return copyAsset(path, targetFolder+string(os.PathSeparator)+targetFilename)
			}

			if useLayout {
				return g.parseFileWithLayout(path, targetFolder+string(os.PathSeparator)+targetFilename, layout)
			}

			return g.parseFile(path, targetFolder+string(os.PathSeparator)+targetFilename)
		}

		targetAssetname := getTargetDirname(srcFolder, path)

		return copyAsset(path, targetFolder+string(os.PathSeparator)+targetAssetname)
	}
}

func (g Generator) parseFile(path, targetFilename string) error {

	dataSource := getSourceFilename(path, g.VarReader.GetDsExtension())

	r, err := os.Open(dataSource)
	if err != nil {
		return err
	}
	defer r.Close()

	m, err := g.extractVariables(r)
	if err != nil {
		return err
	}

	f, err := os.Create(targetFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, m)
	if err != nil {
		return err
	}

	return nil
}

func (g Generator) parseFileWithLayout(path, targetFilename, layout string) error {

	if path == layout {
		return nil
	}

	dataSource := getSourceFilename(path, g.VarReader.GetDsExtension())
	layoutSource := getSourceFilename(layout, g.VarReader.GetDsExtension())

	r, err := mergeSourceFile(layoutSource, dataSource)
	if err != nil {
		return err
	}

	m, err := g.extractVariables(r)
	if err != nil {
		return err
	}

	f, err := os.Create(targetFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl, err := template.ParseFiles(layout, path)
	if err != nil {
		return err
	}

	err = tmpl.ExecuteTemplate(f, g.Config.Base, m)
	if err != nil {
		return err
	}

	return nil

}

func (g Generator) extractVariables(r io.Reader) (map[interface{}]interface{}, error) {
	m, err := g.VarReader.GetVarsForTpl(r)
	if err != nil {
		return nil, err
	}

	return m, nil
}
