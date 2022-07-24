package resources

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var ErrDbNotInitialized = errors.New("database not initialized")

func InitializeSQLite() {
	if db == nil {
		d, err := sql.Open("sqlite3", "stage.db")
		if err != nil {
			panic(err)
		}
		db = d
	}
}

func GetSQLite() *sql.DB {
	return db
}

func CloseSQLite() error {
	if db != nil {
		return db.Close()
	}
	return ErrDbNotInitialized
}
