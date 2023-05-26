package mocks

import (
	"github.com/lelebus/inditext-go-backend/internal/database"
	"github.com/lelebus/inditext-go-backend/internal/domain"
)

var MockedBrands = []*domain.Brand{
	{
		ID:   1,
		Name: "ZARA",
	},
}

func ResetBrands() {
	for _, brand := range MockedBrands {
		if _, err := database.CreateBrand(*brand); err != nil {
			panic(err)
		}
	}
}
