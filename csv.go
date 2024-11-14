package csv

import (
	"bytes"
	"encoding/csv"
	"fmt"

	"github.com/jszwec/csvutil"
)

// Formatter is the CSV formatter for Gout
type Formatter struct {
	Headers []string
}

func (w Formatter) Format(v interface{}) ([]byte, error) {
	// Just marshal if we don't have custom headers
	if w.Headers == nil {
		return csvutil.Marshal(v)
	}
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Write the custom header
	if err := writer.Write(w.Headers); err != nil {
		return nil, fmt.Errorf("failed to write header: %w", err)
	}

	// Marshal data into CSV format
	enc := csvutil.NewEncoder(writer)
	enc.AutoHeader = false
	if err := enc.Encode(v); err != nil {
		return nil, fmt.Errorf("failed to encode data: %w", err)
	}

	// Flush and check for any errors from the writer
	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, fmt.Errorf("error writing CSV: %w", err)
	}

	return buf.Bytes(), nil
}
