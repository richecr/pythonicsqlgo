package dialects

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/richecr/pythonic-core/lib/query"
)

type Client struct {
	Uri      string
	Dialect  string
	Database *sql.DB
	Compiler *query.QueryCompiler
}

func NewClient(dialect string, uri string) *Client {
	client := &Client{
		Uri:      uri,
		Dialect:  dialect,
		Database: nil,
		Compiler: nil,
	}
	client.init()
	return client
}

func (c *Client) init() {
	db, err := sql.Open(c.Dialect, c.Uri)
	if err != nil {
		panic(err)
	}
	c.Database = db
	c.Compiler = query.NewQueryCompiler(db)
}

func (c *Client) Exec(sql string) ([]map[string]interface{}, error) {
	result, err := c.Compiler.Exec(sql)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) Disconnect() error {
	err := c.Database.Close()
	if err != nil {
		return err
	}
	return nil
}
