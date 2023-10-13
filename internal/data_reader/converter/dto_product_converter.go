package converter

import (
	"errors"
	"linxTestCase/internal/dto"
	"strconv"
)

func DtoProductConverter(vals []string) (dto.Product, error) {
	if len(vals) != 3 {
		return dto.Product{}, errors.New("Number of arguments doesn't match Product DTO")
	}
	prod := dto.Product{}
	prod.Name = vals[0]
	var err error
	prod.Price, err = strconv.Atoi(vals[1])
	if err != nil {
		return dto.Product{}, err
	}
	prod.Rating, err = strconv.Atoi(vals[2])
	if err != nil {
		return dto.Product{}, err
	}
	return prod, nil
}
