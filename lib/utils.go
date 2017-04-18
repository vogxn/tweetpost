package lib

import (
	"path/filepath"
)

func LayoutFiles() []string {
	files, err := filepath.Glob("templates/layouts/*.tmpl")
	if err != nil {
		panic(err)
	}
	return files
}
