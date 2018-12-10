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
	db, err := open(database)
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
		err = initSchemaMigrations(db)
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

func open(database string) (db *sqlx.DB, err error) {
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

func insertMigrations(n int, db *sqlx.DB) (err error) {
	sql := `INSERT INTO schema_migrations ("version") VALUES ($1)`
	_, err = db.Exec(sql, n)
	return
}

func schemaMigrationsExists(db *sqlx.DB) (b bool, err error) {
	s := struct {
		Select interface{} `db:"s"`
	}{}
	err = db.Get(&s, "SELECT to_regclass('schema_migrations') AS s")
	b = s.Select != nil
	return
}

func createMigrationTable(db *sqlx.DB) (err error) {
	sql := `CREATE TABLE schema_migrations (version bigint NOT NULL, CONSTRAINT schema_migrations_pkey PRIMARY KEY (version))`
	_, err = db.Exec(sql)
	return
}

func migrationMax(db *sqlx.DB) (m int, err error) {
	s := struct {
		Max int `db:"m"`
	}{}
	err = db.Get(&s, `SELECT max("version") AS m FROM schema_migrations`)
	m = s.Max
	return
}

func initSchemaMigrations(db *sqlx.DB) (err error) {
	var b bool
	b, err = schemaMigrationsExists(db)
	if err != nil {
		return
	}
	if !b {
		err = createMigrationTable(db)
	}
	return
}
