package aztable

import (
	"io"

	"github.com/Azure/azure-sdk-for-go/sdk/monitor/azquery"
)

type reader struct {
	source *azquery.Table
	row    int
	offset int
	logCol int
}

func NewReader(aztable *azquery.Table, logCol int) io.Reader {
	return &reader{
		source: aztable,
		logCol: logCol,
	}
}

func (r *reader) Read(p []byte) (n int, err error) {
	if r.source == nil {
		return 0, io.EOF
	}

	bufLen := len(p)
	var bpos int

	for bpos < bufLen {
		cp, err := r.copyRow(p[bpos:])
		if err != nil {
			return bpos, err
		}
		bpos += cp

	}

	return bpos, nil
}

func (r *reader) copyRow(p []byte) (int, error) {
	rowCount := len(r.source.Rows)
	if r.row >= rowCount {
		return 0, io.EOF
	}

	// Handle both string values and nil values
	var currRow string

	// Read rows in reverse order
	rowValue := r.source.Rows[rowCount-1-r.row][r.logCol]

	// Handle the various possible types
	switch {
	case rowValue == nil:
		// Treat nil as an empty string
		currRow = "\n"
	case rowValue == "":
		// Empty string is already handled correctly
		currRow = "\n"
	default:
		// Try to convert to string
		if strValue, ok := rowValue.(string); ok {
			currRow = strValue + "\n"
		} else {
			// If not a string and not nil, skip this row
			r.row++
			r.offset = 0
			if r.row >= rowCount {
				return 0, io.EOF
			}
			return r.copyRow(p) // Try the next row recursively
		}
	}

	cp := copy(p, currRow[r.offset:])
	r.offset += cp

	if r.offset >= len(currRow) {
		r.row++
		r.offset = 0
	}
	return cp, nil
}
