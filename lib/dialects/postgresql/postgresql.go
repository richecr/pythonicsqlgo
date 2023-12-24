package postgresql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresSQL struct {
	Dialect  string
	Uri      string
	Database *sql.DB
}

func NewPostgresSQL(dialect string, uri string) *PostgresSQL {
	return &PostgresSQL{
		Dialect:  dialect,
		Uri:      uri,
		Database: nil,
	}
}

func (p *PostgresSQL) Connect() error {
	db, err := sql.Open(p.Dialect, p.Uri)
	if err != nil {
		return err
	}
	p.Database = db
	return nil
}

func (p *PostgresSQL) Disconnect() error {
	if p.Database != nil {
		err := p.Database.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
