package pythonic

import (
	"github.com/richecr/pythonicsqlgo/lib/dialects"
	"github.com/richecr/pythonicsqlgo/lib/query"
	"github.com/richecr/pythonicsqlgo/lib/query/model"
)

type PythonicSQL struct {
	Uri     string
	Dialect string
	Client  dialects.Client
	Query   *query.QueryBuilder
}

func NewPythonicSQL(config model.DatabaseConfiguration) (*PythonicSQL, error) {
	sql := &PythonicSQL{
		Uri:     config.Config.Uri,
		Dialect: config.Client,
	}

	client, err := sql.getClient()
	if err != nil {
		return nil, err
	}
	sql.Client = *client
	sql.Query = *&client.Builder

	return sql, nil
}

func (p *PythonicSQL) getClient() (*dialects.Client, error) {
	var client = dialects.NewClient(p.Dialect, p.Uri)

	return client, nil
}
