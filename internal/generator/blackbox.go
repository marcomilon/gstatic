package generator

import (
	"os"
	"path/filepath"
)

// Blackbox is a resolver that does nothing
type Blackbox struct {
	Source string
	Target string
}

func (blackbox Blackbox) resolver() filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return nil
	}
}

func (blackbox Blackbox) source() string {
	return blackbox.Source
}
