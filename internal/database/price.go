package database

import (
	"time"

	"github.com/lelebus/inditext-go-backend/internal/domain"
)

func CreatePrice(price domain.Price) (*domain.Price, error) {
	result := database.Create(&price)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &price, nil
}

func GetPriceByApplicationDate(brandId, productId string, applicationDate time.Time) (price *domain.Price, err error) {
	price = &domain.Price{}
	result := database.Where("brand_id = ? AND product_id = ? AND start_date <= ? AND end_date >= ?", brandId, productId, applicationDate, applicationDate).Order("priority desc").First(&price)
	if err = result.Error; err != nil {
		return nil, err
	}

	return price, nil
}

