package aztable_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/monitor/azquery"
	"github.com/equinor/radix-log-api/pkg/aztable"
	"github.com/stretchr/testify/assert"
)

func TestAzTableReader(t *testing.T) {
	tests := []struct {
		name        string
		tableData   [][]interface{}
		logColIndex int
		expected    string
		expectedErr error
	}{
		{
			name: "reads single row correctly",
			tableData: [][]interface{}{
				{"timestamp", "log message content"},
			},
			logColIndex: 1,
			expected:    "log message content\n",
			expectedErr: nil,
		},
		{
			name: "reads multiple rows in reverse order",
			tableData: [][]interface{}{
				{"timestamp1", "first log message"},
				{"timestamp2", "second log message"},
				{"timestamp3", "third log message"},
			},
			logColIndex: 1,
			expected:    "third log message\nsecond log message\nfirst log message\n",
			expectedErr: nil,
		},
		{
			name: "handles empty log message",
			tableData: [][]interface{}{
				{"timestamp", ""},
			},
			logColIndex: 1,
			expected:    "\n",
			expectedErr: nil,
		},
		{
			name: "handles mixed content with empty lines",
			tableData: [][]interface{}{
				{"timestamp1", "first message"},
				{"timestamp2", ""},
				{"timestamp3", "third message"},
				{"timestamp4", nil},
				{"timestamp5", "fifth message"},
			},
			logColIndex: 1,
			expected:    "fifth message\n\nthird message\n\nfirst message\n",
			expectedErr: nil,
		},
		{
			name:        "handles empty table",
			tableData:   [][]interface{}{},
			logColIndex: 1,
			expected:    "",
			expectedErr: nil,
		},
		{
			name: "handles content with unknown data types",
			tableData: [][]interface{}{
				{"timestamp1", "first message"},
				{"timestamp2", 1234},
				{"timestamp3", "third message"},
			},
			logColIndex: 1,
			expected:    "third message\n",
			expectedErr: aztable.ErrUnexpectedData,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create an Azure Table with the test data
			var rows []azquery.Row
			for _, rowData := range tt.tableData {
				rows = append(rows, rowData)
			}

			table := &azquery.Table{
				Rows: rows,
			}

			// Create a new reader
			reader := aztable.NewReader(table, tt.logColIndex)

			// Read all content
			buf := new(bytes.Buffer)
			_, err := io.Copy(buf, reader) // Ignoring error as io.Copy always returns EOF at the end
			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, buf.String())
		})
	}
}

func TestAzTableReaderWithChunkedReading(t *testing.T) {
	// Setup test data
	tableData := [][]interface{}{
		{"timestamp1", "first line"},
		{"timestamp2", "second line that is longer"},
		{"timestamp3", "third\nwith newline"},
	}

	var rows []azquery.Row
	for _, rowData := range tableData {
		rows = append(rows, rowData)
	}

	table := &azquery.Table{
		Rows: rows,
	}

	// Create reader
	reader := aztable.NewReader(table, 1)

	// Test reading in small chunks
	buf := new(bytes.Buffer)

	// Read in chunks of 5 bytes
	chunk := make([]byte, 5)
	for {
		n, err := reader.Read(chunk)
		buf.Write(chunk[:n])
		if err == io.EOF {
			break
		}
		assert.NoError(t, err)
	}

	// Expected output (in reverse order)
	expected := "third\nwith newline\nsecond line that is longer\nfirst line\n"
	assert.Equal(t, expected, buf.String())
}

func TestReaderWithNilSource(t *testing.T) {
	// Test with nil source
	reader := aztable.NewReader(nil, 0)

	buf := make([]byte, 10)
	n, err := reader.Read(buf)

	assert.Equal(t, 0, n)
	assert.Equal(t, io.EOF, err)
}

func TestInvalidColumnData(t *testing.T) {
	// Test with non-string data in the log column
	tableData := [][]interface{}{
		{"timestamp", 123}, // Integer instead of string
	}

	var rows []azquery.Row
	for _, rowData := range tableData {
		rows = append(rows, rowData)
	}

	table := &azquery.Table{
		Rows: rows,
	}

	reader := aztable.NewReader(table, 1)

	buf := make([]byte, 10)
	n, err := reader.Read(buf)

	// The reader should return an error with message "unexpected data in log"
	// when encountering non-string values in the log column.
	assert.Equal(t, 0, n)
	assert.Equal(t, "unexpected data in log", err.Error())
}
