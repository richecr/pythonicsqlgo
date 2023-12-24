package model

type SimpleAttributes struct {
	TableName string
	Limit     int
	Offset    int
	Counter   int
	IsDQL     bool
	Raw       Raw
}

func NewSimpleAttributes(tableName string, limit, offset, counter int, isDQL bool) *SimpleAttributes {
	return &SimpleAttributes{
		TableName: tableName,
		Limit:     limit,
		Offset:    offset,
		Counter:   counter,
		IsDQL:     isDQL,
	}
}
