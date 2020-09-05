package generator

import (
	"os"
	"path/filepath"
	"text/template"
)

// Generator is the main struct. It use a VarReader to extract variables from a data source
type Generator struct {
	VarReader VarReader
}

// Generate is the main method used to generate a site
func (g Generator) Generate(srcFolder string, targetFolder string) error {

	err := validateTargetFolder(targetFolder)
	if err != nil {
		return err
	}

	resolver := g.resolver(srcFolder, targetFolder)

	return filepath.Walk(srcFolder, resolver)
}

func (g Generator) resolver(srcFolder string, targetFolder string) filepath.WalkFunc {

	var useLayout bool = false
	var layout string = srcFolder + string(os.PathSeparator) + "layout" + string(os.PathSeparator) + "layout.html"

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
		if ext == ".html" {
			targetFilename := getTargetDirname(srcFolder, path)

			sourceFile := hasSourceFilename(path)
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

	m, err := g.extractVariables(path)
	if err != nil {
		return err
	}

	f, err := os.Create(targetFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl, err := template.ParseFiles(path)

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

	m, err := g.extractVariables(path)
	if err != nil {
		return err
	}

	f, err := os.Create(targetFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl, err := template.ParseFiles(layout, path)
	err = tmpl.ExecuteTemplate(f, "base", m)
	if err != nil {
		return err
	}

	return nil

}

func (g Generator) extractVariables(path string) (map[interface{}]interface{}, error) {
	dataSource := getSourceFilename(path)

	r, err := os.Open(dataSource)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	m, err := g.VarReader.GetVarsForTpl(r)
	if err != nil {
		return nil, err
	}

	return m, nil
}
