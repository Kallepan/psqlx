package psqlx_test

import (
	"database/sql"
	"testing"

	_ "github.com/kallepan/psqlx"
)

func TestInit(t *testing.T) {
	db, err := sql.Open("psqlx", "postgres://postgres:postgres@db:5432/postgres?sslmode=disable")
	if err != nil {
		t.Errorf("Error opening database: %s", err.Error())
	}

	if err := db.Ping(); err != nil {
		t.Errorf("Error pinging database: %s", err.Error())
	}
}
