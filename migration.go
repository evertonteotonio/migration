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

// Run parse and performs the required migration
func Run(source, database, migrate string) (err error) {
	var n int
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
		err = up(source, n)
	case "down":
		n, err = parsePar(m)
		if err != nil {
			return
		}
		err = down(source, n)
	default:
		err = fmt.Errorf("unknown migration command")
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
