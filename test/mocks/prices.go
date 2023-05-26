package mocks

import (
	"time"

	"github.com/lelebus/inditext-go-backend/internal/database"
	"github.com/lelebus/inditext-go-backend/internal/domain"
)

var MockedPrices = []*domain.Price{
	{
		BrandID:     1,
		ProductID:   35455,
		PriceListID: 1,
		StartDate:   parseDateTime("2020-06-14-00.00.00"),
		EndDate:     parseDateTime("2020-12-31-23.59.59"),
		Priority:    0,
		Price:       3550,
		Currency:    "EUR",
	},
	{
		BrandID:     1,
		ProductID:   35455,
		PriceListID: 2,
		StartDate:   parseDateTime("2020-06-14-15.00.00"),
		EndDate:     parseDateTime("2020-06-14-18.30.00"),
		Priority:    1,
		Price:       2545,
		Currency:    "EUR",
	},
	{
		BrandID:     1,
		ProductID:   35455,
		PriceListID: 3,
		StartDate:   parseDateTime("2020-06-15-00.00.00"),
		EndDate:     parseDateTime("2020-06-15-11.00.00"),
		Priority:    1,
		Price:       3050,
		Currency:    "EUR",
	},
	{
		BrandID:     1,
		ProductID:   35455,
		PriceListID: 4,
		StartDate:   parseDateTime("2020-06-15-16.00.00"),
		EndDate:     parseDateTime("2020-12-31-23.59.59"),
		Priority:    1,
		Price:       3895,
		Currency:    "EUR",
	},
}

func parseDateTime(dateTimeStr string) time.Time {
	layout := "2006-01-02-15.04.05"
	t, err := time.Parse(layout, dateTimeStr)
	if err != nil {
		panic(err)
	}
	return t
}

func ResetPrices() {
	for _, price := range MockedPrices {
		if _, err := database.CreatePrice(*price); err != nil {
			panic(err)
		}
	}
}
