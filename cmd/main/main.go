package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"linxTestCase/internal/data_reader"
	"linxTestCase/internal/data_reader/converter"
	"linxTestCase/internal/dto"
	"linxTestCase/internal/model"
	"linxTestCase/internal/usecase"
	"os"
	"path"
	"strings"
)

func main() {
	// reading cmd arg
	envFPath, _ := os.LookupEnv("FILE_PATH")
	flagFPath := flag.String("f", "", "Path to file from which to extract data")
	flag.Parse()

	fPath := envFPath
	if *flagFPath != "" {
		fPath = *flagFPath
	}
	fPath = strings.TrimSpace(fPath)

	fmt.Printf("Reading from \"%s\"\n", fPath)

	// opening file
	f, err := os.Open(fPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer f.Close()
	rd := bufio.NewReader(f)

	// creating appropriate reader
	var reader data_reader.ObjectReader[dto.Product]
	switch strings.ToLower(path.Ext(fPath)) {
	case ".json":
		reader = setUpJsonReader[dto.Product](rd)
	case ".csv":
		reader = setUpCsvReader[dto.Product](rd, converter.DtoProductConverter)
	default:
		fmt.Printf("File's extension does not match .csv or .json")
		os.Exit(-1)
	}

	// going through data
	var maxPriceProd *model.Product
	var maxRatingProd *model.Product
	maxPriceProd, maxRatingProd, err = usecase.FindMaxProduct(reader)

	// printing the results
	if maxPriceProd != nil {
		fmt.Printf("Product with max price: %s with price of %d\n", maxPriceProd.Name, maxPriceProd.Price)
		fmt.Printf("Product with max rating: %s with rating of %d", maxRatingProd.Name, maxRatingProd.Rating)
	} else {
		fmt.Printf("No data provided")
	}
}

func setUpJsonReader[K any](rd io.Reader) data_reader.ObjectReader[K] {
	return data_reader.NewJSONReader[K](rd)
}

func setUpCsvReader[K any](rd io.Reader, transformer func([]string) (K, error)) data_reader.ObjectReader[K] {
	reader := data_reader.NewCsvReader[K](rd, transformer)
	reader.Delimer(',')
	reader.SkipHeader(true)
	return reader
}
