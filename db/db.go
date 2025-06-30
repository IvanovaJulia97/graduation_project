package db

import (
	"database/sql"
)

var db *sql.DB

const schema = `
CREATE TABLE IF NOT EXISTS scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date CHAR(8) NOT NULL DEFAULT '',
	title VARCHAR(32) NOT NULL DEFAULT '',
	comment TEXT,
	repeat VARCHAR(128) DEFAULT ''
);

CREATE INDEX IF NOT EXISTS idx_scheduler_date ON scheduler(date);
`

func Init(dbFile string) error {
	connect, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	db = connect

	// ВСЕГДА выполняем schema, потому что она безопасна (IF NOT EXISTS)
	if _, err := db.Exec(schema); err != nil {
		db.Close()
		return err
	}

	return nil
}

func Get() *sql.DB {
	return db
}
