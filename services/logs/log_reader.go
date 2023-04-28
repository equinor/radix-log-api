package logs

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/monitor/azquery"
)

type logReader struct {
	source *azquery.Table
	row    int
	offset int
	logCol int
}

func mustParseTime(t string) time.Time {
	if t == "" {
		fmt.Println("")
	}
	parsed, err := time.Parse(time.RFC3339, t)
	if err != nil {
		panic(err)
	}
	return parsed
}

func (r *logReader) Read(p []byte) (n int, err error) {
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

func (r *logReader) copyRow(p []byte) (int, error) {
	rowCount := len(r.source.Rows)
	if r.row >= rowCount {
		return 0, io.EOF
	}
	currRow, ok := r.source.Rows[r.row][r.logCol].(string)
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
