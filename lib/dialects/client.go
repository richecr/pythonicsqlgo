package dialects

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/richecr/pythonicsqlgo/lib/query"
)

type Client struct {
	Uri      string
	Dialect  string
	Database *sql.DB
	Compiler *query.QueryCompiler
	Builder  *query.QueryBuilder
}

func NewClient(dialect string, uri string) *Client {
	client := &Client{
		Uri:      uri,
		Dialect:  dialect,
		Database: nil,
		Compiler: nil,
		Builder:  nil,
	}
	client.init()
	return client
}

func (c *Client) init() {
	db, err := sql.Open(c.Dialect, c.Uri)
	if err != nil {
		fmt.Println("Ops! Ocorreu um erro ao abrir a conexão com o banco de dados.")
	}
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
	c.Database = db
	c.Compiler = query.NewQueryCompiler(db)
	c.Builder = query.NewQueryBuilder(c.Compiler)
}

func (c *Client) Connect() error {
	err := c.Database.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Disconnect() error {
	err := c.Database.Close()
	if err != nil {
		return err
	}
	return nil
}
