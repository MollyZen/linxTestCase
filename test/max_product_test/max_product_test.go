package max_product_test

import (
	"linxTestCase/internal/data_reader"
	"linxTestCase/internal/dto"
	"linxTestCase/internal/usecase"
	"testing"
)

func TestDefaultMaxProduct1(t *testing.T) {
	vals := []dto.Product{
		dto.Product{"Печенье", 34, 3},
		dto.Product{"Сахар", 45, 2},
		dto.Product{"Варенье", 200, 5},
	}
	rd := data_reader.NewArrayObjectReader(vals)
	r1, r2, err := usecase.FindMaxProduct(rd)
	if err != nil {
		t.Fatal(err)
	}
	if r1 == nil || r1.Name != vals[2].Name {
		t.Fatal("Wrong result with max price")
	}
	if r1 == nil || r2.Name != vals[2].Name {
		t.Fatal("Wrong result with max rating")
	}
}

func TestDefaultMaxProduct2(t *testing.T) {
	vals := []dto.Product{
		dto.Product{"Печенье", 3, 5},
		dto.Product{"Яблоки", 1, 2},
		dto.Product{"Тыква", 2, 3},
	}
	rd := data_reader.NewArrayObjectReader(vals)
	r1, r2, err := usecase.FindMaxProduct(rd)
	if err != nil {
		t.Fatal(err)
	}
	if r1 == nil || r1.Name != vals[0].Name {
		t.Fatal("Wrong result with max price")
	}
	if r2 == nil || r2.Name != vals[0].Name {
		t.Fatal("Wrong result with max rating")
	}
}

func TestDifferentValuesMaxProduct(t *testing.T) {
	vals := []dto.Product{
		dto.Product{"Печенье", 3, 3},
		dto.Product{"Яблоки", 1, 2},
		dto.Product{"Тыква", 2, 5},
	}
	rd := data_reader.NewArrayObjectReader(vals)
	r1, r2, err := usecase.FindMaxProduct(rd)
	if err != nil {
		t.Fatal(err)
	}
	if r1 == nil || r1.Name != vals[0].Name {
		t.Fatal("Wrong result with max price")
	}
	if r2 == nil || r2.Name != vals[2].Name {
		t.Fatal("Wrong result with max rating")
	}
}

func TestDifferentMaxObjectsMaxProduct(t *testing.T) {
	vals := []dto.Product{
		dto.Product{"Печенье", 3, 3},
		dto.Product{"Яблоки", 1, 2},
		dto.Product{"Тыква", 2, 5},
	}
	rd := data_reader.NewArrayObjectReader(vals)
	r1, r2, err := usecase.FindMaxProduct(rd)
	if err != nil {
		t.Fatal(err)
	}
	if r1 == nil || r1.Name != vals[0].Name {
		t.Fatal("Wrong result with max price")
	}
	if r2 == nil || r2.Name != vals[2].Name {
		t.Fatal("Wrong result with max rating")
	}
}

func TestOneValueMaxProduct(t *testing.T) {
	vals := []dto.Product{
		dto.Product{"Печенье", 3, 3},
	}
	rd := data_reader.NewArrayObjectReader(vals)
	r1, r2, err := usecase.FindMaxProduct(rd)
	if err != nil {
		t.Fatal(err)
	}
	if r1 == nil || r1.Name != vals[0].Name {
		t.Fatal("Wrong result with max price")
	}
	if r2 == nil || r2.Name != vals[0].Name {
		t.Fatal("Wrong result with max rating")
	}
}

func TestNoValuesMaxProduct(t *testing.T) {
	vals := []dto.Product{}
	rd := data_reader.NewArrayObjectReader(vals)
	r1, r2, err := usecase.FindMaxProduct(rd)
	if err != nil {
		t.Fatal(err)
	}
	if r1 != nil {
		t.Fatal("Wrong result with max price")
	}
	if r2 != nil {
		t.Fatal("Wrong result with max rating")
	}
}
