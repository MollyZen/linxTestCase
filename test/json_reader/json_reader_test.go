package json_reader

import (
	"bufio"
	"fmt"
	"io"
	"linxTestCase/internal/data_reader"
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

func TestFullyEmptyJsonReader(t *testing.T) {
	f, rd, err := openFileReader("fully_empty.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	jsonRd := data_reader.NewJSONReader[dto.Product](rd)

	for jsonRd.More() {
		_, err = jsonRd.ReadNext()
		if err != nil {
			t.Fatal(err)
		}
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestOnlyBracketsJsonReader(t *testing.T) {
	f, rd, err := openFileReader("only_brackets.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	jsonRd := data_reader.NewJSONReader[dto.Product](rd)

	for jsonRd.More() {
		_, err = jsonRd.ReadNext()
		if err != nil {
			t.Fatal(err)
		}
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestExampleValueJsonReader(t *testing.T) {
	f, rd, err := openFileReader("db.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	jsonRd := data_reader.NewJSONReader[dto.Product](rd)

	for jsonRd.More() {
		_, err = jsonRd.ReadNext()
		if err != nil {
			t.Fatal(err)
		}
	}
	if err != nil {
		t.Fatal(err)
	}
}
