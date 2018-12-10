package migration

import (
	"fmt"
	"path/filepath"
	"sort"
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
	sort.Sort(sort.Reverse(sort.StringSlice(files)))
	return
}

func up(source string, start, n int) (err error) {
	files, err := upFiles(source)
	if err != nil {
		return
	}
	err = exec(files, start, n)
	return
}

func down(source string, start, n int) (err error) {
	files, err := downFiles(source)
	if err != nil {
		return
	}
	err = exec(files, start, n)
	return
}

func exec(files []string, start, n int) (err error) {
	if n == 0 {
		n = len(files)
	}
	for _, f := range files[start:n] {
		fmt.Println(f)
	}
	return
}

func parsePar(m []string) (n int, err error) {
	if len(m) > 1 {
		n, err = strconv.Atoi(m[1])
		if err != nil {
			err = fmt.Errorf("invalid syntax")
			return
		}
	}
	return
}

// Run parse and performs the required migration
func Run(source, database, migrate string) (err error) {
	var start, n int
	m := strings.Split(migrate, " ")
	if len(m) > 2 {
		err = fmt.Errorf("the number of migration parameters is incorrect")
		return
	}
	switch m[0] {
	case "up":
		n, err = parsePar(m)
		if err != nil {
			return
		}
		err = up(source, start, n)
	case "down":
		n, err = parsePar(m)
		if err != nil {
			return
		}
		err = down(source, start, n)
	default:
		err = fmt.Errorf("unknown migration command")
	}
	return
}
