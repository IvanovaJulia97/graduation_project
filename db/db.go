package db

import (
	"database/sql"
)

var DB *sql.DB

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

	DB = connect

	if _, err := DB.Exec(schema); err != nil {
		DB.Close()
		return err
	}

	return nil
}

func Get() *sql.DB {
	return DB
}
