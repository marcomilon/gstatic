package generator

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/marcomilon/gstatic/internal/helpers"
)

// Generator is the main struct. It use a VarReader to extract variables from a data source
type Generator struct {
	VarReader VarReader
}

// Generate is the main method used to generate a site
func (g Generator) Generate(source string, target string) error {
	resolver := g.resolver(source, target)

	err := filepath.Walk(source, resolver)
	if err != nil {
		log.Println(err)
	}

	return err
}

func (g Generator) resolver(source string, target string) filepath.WalkFunc {

	layouts, err := helpers.ReadLayouts(source + layoutFolder)
	if err != nil {
		log.Printf("Warning: layouts not found, %v", err)
	}

	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {

			if filepath.Ext(path) != templateExt {
				return nil
			}

			if err := g.parseFile(layouts, path, source, target); err != nil {
				return err
			}

			return nil
		}

		targetFolder := target + removeSourceFromTemplatePath(path, source)
		if err := helpers.CreateDir(targetFolder); err != nil {
			return err
		}

		return nil
	}
}

func (g Generator) parseFile(layouts []string, path, source, target string) error {

	if filepath.Dir(path) == source+layoutFolder {
		return nil
	}

	var html []string

	m, err := g.VarReader.GetVarsForTpl(path)
	if err != nil {
		return err
	}

	templateSrc := removeSourceFromTemplatePath(path, source)
	targetTemplate := target + templateSrc

	f, err := os.Create(targetTemplate)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return err
	}
	defer f.Close()

	if len(layouts) > 0 {
		html = append(layouts, path)
	} else {
		html = append(html, path)
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

func removeSourceFromTemplatePath(path string, source string) string {
	return strings.Replace(path, source, "", 1)
}
