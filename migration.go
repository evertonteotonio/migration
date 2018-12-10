package migration

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

func up(source string, start, n int, db *sqlx.DB) (err error) {
	files, err := upFiles(source)
	if err != nil {
		return
	}
	err = exec(files, start, n, db)
	return
}

func down(source string, start, n int, db *sqlx.DB) (err error) {
	files, err := downFiles(source)
	if err != nil {
		return
	}
	err = exec(files, start, n, db)
	return
}

func exec(files []string, start, n int, db *sqlx.DB) (err error) {
	if n == 0 {
		n = len(files)
	}
	for _, f := range files[start:n] {
		var b []byte
		b, err = ioutil.ReadFile(f) // nolint
		if err != nil {
			return
		}
		_, err = db.Exec(string(b))
		if err != nil {
			return
		}
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

	//"postgres://postgres@localhost:5432/cesar?sslmode=disable")
	db, err := Open(database)
	if err != nil {
		return
	}

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
		err = up(source, start, n, db)
	case "down":
		n, err = parsePar(m)
		if err != nil {
			return
		}
		err = down(source, start, n, db)
	default:
		err = fmt.Errorf("unknown migration command")
	}
	return
}

func Open(database string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open("postgres", database)
	if err != nil {
		err = fmt.Errorf("error open db: %v", err)
		return
	}
	err = db.Ping()
	if err != nil {
		err = fmt.Errorf("error ping db: %v", err)
	}
	return
}
