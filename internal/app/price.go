package app

import (
	"errors"
	"time"

	"github.com/lelebus/inditext-go-backend/internal/database"
	"github.com/lelebus/inditext-go-backend/internal/domain"

	"github.com/jinzhu/gorm"
)

func GetPriceByApplicationDate(brandId, productId string, applicationDate time.Time) (*domain.Price, error) {
	price, err := database.GetPriceByApplicationDate(brandId, productId, applicationDate)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return price, nil
}
