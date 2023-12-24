package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	Dialect  string
	Uri      string
	Database *sql.DB
}

func NewSQLite(dialect string, uri string) *SQLite {
	return &SQLite{
		Dialect:  dialect,
		Uri:      uri,
		Database: nil,
	}
}

func (s *SQLite) Connect() error {
	db, err := sql.Open("sqlite3", s.Uri)
	if err != nil {
		return err
	}
	s.Database = db
	return nil
}

func (s *SQLite) Disconnect() error {
	if s.Database != nil {
		err := s.Database.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
