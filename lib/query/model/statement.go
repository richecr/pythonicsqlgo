package model

type Statements struct {
	Typ       string
	Value     interface{}
	Grouping  string
	Column    string
	Condition string
	Operator  string
}

// func NewStatements(typeVar, grouping, column, condition, operator string) *Statements {
// 	return &Statements{
// 		TypeVar:   typeVar,
// 		Value:     value,
// 		Grouping:  grouping,
// 		Column:    column,
// 		Condition: condition,
// 		Operator:  operator,
// 	}
// }
