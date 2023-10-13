package csv_reader

import (
	"bufio"
	"fmt"
	"io"
	"linxTestCase/internal/data_reader"
	"linxTestCase/internal/data_reader/converter"
	"linxTestCase/internal/dto"
	"os"
	"testing"
)

func openFileReader(path string) (*os.File, io.Reader, error) {
	fName := path
	f, err := os.Open(fName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return f, bufio.NewReader(f), nil
}

func TestFullyEmptyCsvReader(t *testing.T) {
	f, rd, err := openFileReader("fully_empty.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	csvRd := data_reader.NewCsvReader[dto.Product](rd, converter.DtoProductConverter)
	csvRd.Delimer(',')
	csvRd.SkipHeader(true)

	for csvRd.More() {
		_, err = csvRd.ReadNext()
		if err != nil {
			t.Fatal(err)
		}
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestOnlyHeaderCsvReader(t *testing.T) {
	f, rd, err := openFileReader("only_header.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	csvRd := data_reader.NewCsvReader[dto.Product](rd, converter.DtoProductConverter)
	csvRd.Delimer(',')
	csvRd.SkipHeader(true)

	for csvRd.More() {
		_, err = csvRd.ReadNext()
		if err != nil {
			t.Fatal(err)
		}
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestExampleValueCsvReader(t *testing.T) {
	f, rd, err := openFileReader("db.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	csvRd := data_reader.NewCsvReader[dto.Product](rd, converter.DtoProductConverter)
	csvRd.Delimer(',')
	csvRd.SkipHeader(true)

	for csvRd.More() {
		_, err = csvRd.ReadNext()
		if err != nil {
			t.Fatal(err)
		}
	}
	if err != nil {
		t.Fatal(err)
	}
}
