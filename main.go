package main

import (
	"fmt"
	"time"

	"github.com/richecr/pythonicsqlgo/lib/pythonic"
	"github.com/richecr/pythonicsqlgo/lib/query/model"
)

func main() {
	s := time.Now()
	p, err := pythonic.NewPythonicSQL(
		model.DatabaseConfiguration{
			Client: "postgres",
			Config: model.Config{
				Uri: "postgres://dev-duckorm:postgres123@localhost:5432/dev-duckorm?sslmode=disable",
			},
		},
	)
	if err != nil {
		panic(err)
	}
	a, err := p.Query.Select([]string{"id"}).From_("users").Exec()
	for _, row := range a {
		fmt.Println(row)
	}
	// fmt.Println(p.Query.Select([]string{"id", "name"}).From_("users").Exec())
	// fmt.Println(p.Query.Select([]string{"id", "name"}).From_("users").Where("id", "0", "=").Exec())
	// fmt.Println(p.Query.Select([]string{"id", "name"}).From_("users").WhereIn("id", []string{"1", "0"}).Exec())
	// fmt.Println(p.Query.Select([]string{"id", "name"}).From_("users").WhereLike("name", "%Ri%").Exec())
	duration := time.Since(s)
	fmt.Println(duration)
}
