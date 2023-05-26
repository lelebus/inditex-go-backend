package mocks

import "github.com/lelebus/inditext-go-backend/internal/database"

func ResetDB() {
	// delete all records
	if err := database.GetInstance().Exec("DELETE FROM brands").Error; err != nil {
		panic(err)
	}
	if err := database.GetInstance().Exec("DELETE FROM prices").Error; err != nil {
		panic(err)
	}

	// insert brands
	ResetBrands()

	// insert prices
	ResetPrices()
}
