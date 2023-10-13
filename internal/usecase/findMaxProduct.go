package usecase

import (
	"linxTestCase/internal/data_reader"
	"linxTestCase/internal/dto"
	"linxTestCase/internal/model"
)

// FindMaxProduct first returned argument is product with Max Price, Second - with Max Rating
func FindMaxProduct(reader data_reader.ObjectReader[dto.Product]) (*model.Product, *model.Product, error) {
	var maxPriceProd dto.Product
	var maxRatingProd dto.Product
	maxPrice := -1
	maxRating := -1

	// going through data
	for reader.More() {
		tmp, err := reader.ReadNext()
		if err != nil {
			return nil, nil, err
		}
		if tmp.Price > maxPrice {
			maxPriceProd = tmp
			maxPrice = tmp.Price
		}
		if tmp.Rating > maxRating {
			maxRatingProd = tmp
			maxRating = tmp.Rating
		}
	}

	if maxPrice != -1 {
		return &model.Product{Name: maxPriceProd.Name, Price: maxPriceProd.Price, Rating: maxPriceProd.Rating},
			&model.Product{Name: maxRatingProd.Name, Price: maxRatingProd.Price, Rating: maxRatingProd.Rating},
			nil
	} else {
		return nil, nil, nil
	}
}
