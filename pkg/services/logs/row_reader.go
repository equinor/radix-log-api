package logs

import "github.com/Azure/azure-sdk-for-go/sdk/monitor/azquery"

type tableReader map[string]int

func newTableReader(table *azquery.Table) tableReader {
	r := tableReader{}
	for i, col := range table.Columns {
		r[*col.Name] = i
	}
	return r
}

func (r tableReader) Value(row azquery.Row, columnName string, defaultValue any) any {
	col, ok := r[columnName]
	if !ok || col > len(row) {
		return defaultValue
	}
	return row[col]
}
