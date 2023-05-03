package aztable

import (
	"errors"
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
	// Read rows in reverse order
	currRow, ok := r.source.Rows[rowCount-1-r.row][r.logCol].(string)
	if !ok {
		return 0, errors.New("unexpected data in log")
	}
	currRow += "\n"
	cp := copy(p, currRow[r.offset:])
	r.offset += cp

	if r.offset >= len(currRow) {
		r.row++
		r.offset = 0
	}
	return cp, nil
}
