package generator

import (
	"log"
	"path/filepath"
)

type source string

type generator interface {
	resolver() filepath.WalkFunc
	source() string
	target() string
}

// Generate is used to generate a site
func Generate(g generator) error {
	resolver := g.resolver()

	err := filepath.Walk(g.source(), resolver)
	if err != nil {
		log.Println(err)
	}

	return err
}
