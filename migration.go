package migration

import (
	"fmt"
	"path/filepath"
)

// upFiles search for migration up files and return
// a sorted array with the path of all found files
func upFiles(dir string) (files []string, err error) {
	files, err = filepath.Glob(filepath.Join(dir, "*.up.sql"))
	return
}

// downFiles search for migration down files and return
// a sorted array with the path of all found files
func downFiles(dir string) (files []string, err error) {
	files, err = filepath.Glob(filepath.Join(dir, "*.down.sql"))
	return
}

func Run(source, database, migrate string) (err error) {
	files, err := upFiles(source)
	if err != nil {
		return
	}
	for _, f := range files {
		fmt.Println(f)
	}
	return
}
