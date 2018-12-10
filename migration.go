package migration

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
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

func up(source string, n int) (err error) {
	fmt.Println("up", n)

	files, err := upFiles(source)
	if err != nil {
		return
	}
	for _, f := range files {
		fmt.Println(f)
	}
	return
}

func down(source string, n int) (err error) {
	fmt.Println("down", n)
	return
}

func Run(source, database, migrate string) (err error) {
	m := strings.Split(migrate, " ")
	switch m[0] {
	case "up":
		var n int
		n, err = strconv.Atoi(m[1])
		if err != nil {
			return
		}
		err = up(source, n)
	case "down":
		var n int
		n, err = strconv.Atoi(m[1])
		if err != nil {
			return
		}
		err = down(source, n)
	default:
		err = fmt.Errorf("unknown migration command")
	}
	return
}
