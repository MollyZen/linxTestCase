package data_reader

import (
	"encoding/csv"
	"io"
)

type CsvReader[K any] struct {
	rd          *csv.Reader
	transformer func([]string) (K, error)
	tmp         *K
	tmpError    error
	skipHeader  bool
	firstRead   bool
}

func NewCsvReader[K any](rd io.Reader, transformer func([]string) (K, error)) *CsvReader[K] {
	reader := &CsvReader[K]{csv.NewReader(rd), transformer, nil, nil, false, true}
	reader.rd.ReuseRecord = true
	reader.rd.TrimLeadingSpace = true
	return reader
}

func (c *CsvReader[K]) ReadNext() (K, error) {
	var errRes K

	if c.skipHeader && c.firstRead {
		_, _ = c.rd.Read()
	}
	c.firstRead = false

	if c.tmp != nil || c.tmpError != nil {
		tmp := c.tmp
		tmpError := c.tmpError
		c.tmp = nil
		c.tmpError = nil
		return *tmp, tmpError
	}

	record, err := c.rd.Read()
	if err != nil {
		return errRes, err
	}

	res, err := c.transformer(record)
	return res, err
}

func (c *CsvReader[K]) More() bool {
	if c.skipHeader && c.firstRead {
		_, _ = c.rd.Read()
		c.firstRead = false
	}
	record, err := c.rd.Read()
	if err == io.EOF {
		return false
	} else {
		var tmp K
		tmp, c.tmpError = c.transformer(record)
		c.tmp = &tmp
		return true
	}
}

func (c *CsvReader[K]) Delimer(del rune) {
	c.rd.Comma = del
}

func (c *CsvReader[K]) SkipHeader(skip bool) {
	c.skipHeader = skip
}
