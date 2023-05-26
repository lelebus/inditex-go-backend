package app

import (
	"errors"

	"github.com/lelebus/inditext-go-backend/internal/database"
	"github.com/lelebus/inditext-go-backend/internal/domain"

	"github.com/jinzhu/gorm"
)

func GetBrandByID(id string) (*domain.Brand, error) {
	brand, err := database.GetBrandByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return brand, nil
}
