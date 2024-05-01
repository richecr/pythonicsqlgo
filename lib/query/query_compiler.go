package query

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type QueryCompiler struct {
	Client     *sql.DB
}

func NewQueryCompiler(client *sql.DB) *QueryCompiler {
	compiler := &QueryCompiler{
		Client:     client,
	}
	return compiler
}

func (qc *QueryCompiler) Exec(sql string) ([]byte, error) {
	var result []map[string]interface{}

	rows, err := qc.Client.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	vals := make([]interface{}, len(cols))
	for i, ct := range colTypes {
		switch ct.DatabaseTypeName() {
		case "VARCHAR", "TEXT":
			vals[i] = new(string)
		case "INT":
			vals[i] = new(int)
		default:
			vals[i] = new(interface{})
		}
	}

	for rows.Next() {
		scanArgs := make([]interface{}, len(cols))
		for i := range vals {
			scanArgs[i] = &vals[i]
		}

		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println(err)
			continue
		}

		rowMap := make(map[string]interface{})
		for i, colName := range cols {
			valPtr := vals[i]

			var val interface{}
			switch v := valPtr.(type) {
			case *string:
				if v != nil {
					val = *v
				}
			case *int:
				if v != nil {
					val = *v
				}
			default:
				val = v
			}

			rowMap[colName] = val
		}

		result = append(result, rowMap)
	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return jsonResult, nil
}