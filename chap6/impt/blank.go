package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func ExampleAllRows_tableStruct() {
	db, err := sql.Open("sqlite", ":memory:")
	_, _ = db, err
}
