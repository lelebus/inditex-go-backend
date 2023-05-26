package database

import (
	"github.com/lelebus/inditext-go-backend/internal/domain"
)

func CreateBrand(brand domain.Brand) (*domain.Brand, error) {
	result := database.Create(&brand)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &brand, nil
}

func GetBrandByID(id string) (brand *domain.Brand, err error) {
	brand = &domain.Brand{}
	result := database.First(&brand, "id = ?", id)
	if err = result.Error; err != nil {
		return nil, err
	}

	return brand, nil
}
