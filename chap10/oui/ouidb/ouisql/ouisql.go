package ouisql

import (
	"bufio"
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func Generate(filename string) error {
	db, err := sql.Open("sqlite3",
		(&url.URL{Scheme: "file",
			Path:     filepath.ToSlash(filename),
			RawQuery: "mode=rwc&_mutex=no",
			OmitHost: true,
		}).String())
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS oui (id PRIMARY KEY, company NOT NULL)`); err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// OR REPLACE: 由于某些原因, IEEE OUI 存在 id 重复的情况.
	insert, err := tx.Prepare(`INSERT OR REPLACE INTO oui (id, company) VALUES(?,?)`)
	if err != nil {
		return err
	}

	resp, err := http.Get("https://standards-oui.ieee.org/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	lineRegexp := regexp.MustCompile(`([0-9A-F]{6})\s+\(base 16\)\s+(.+)`)
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return err
		}
		m := lineRegexp.FindSubmatch(scanner.Bytes())
		if m == nil {
			continue
		}
		id, err := strconv.ParseUint(string(m[1]), 16, 32)
		if err != nil {
			return fmt.Errorf("invalid id %v: %w", string(m[1]), err)
		}
		if id > 0xFFFFFF {
			return fmt.Errorf("id %v is too large", id)
		}
		company := string(m[2])
		if _, err = insert.Exec(id, company); err != nil {
			return fmt.Errorf("invalid data %v %v: %w", id, company, err)
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return db.Close()
}

type DB struct {
	db    *sql.DB
	query *sql.Stmt
}

func NewDB(filename string) (*DB, error) {
	db, err := sql.Open("sqlite3",
		(&url.URL{
			Scheme:   "file",
			Path:     filepath.ToSlash(filename),
			RawQuery: "mode=ro&_mutex=no",
			OmitHost: true,
		}).String())
	if err != nil {
		return nil, err
	}
	query, err := db.Prepare(`SELECT company FROM oui WHERE id=?`)
	if err != nil {
		return nil, err
	}
	return &DB{db: db, query: query}, nil
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) Lookup(oui string) (company string, err error) {
	id, err := strconv.ParseUint(oui, 16, 32)
	if err != nil {
		err = fmt.Errorf("invalid id %v: %w", oui, err)
		return
	}
	if id > 0xFFFFFF {
		err = fmt.Errorf("id %v is too large", id)
		return
	}
	r := db.query.QueryRow(id)
	if err = r.Scan(&company); err == sql.ErrNoRows {
		err = nil
	}
	return
}
