package domain

import "time"

type Price struct {
	BrandID     uint      `json:"brand_id" gorm:"not null; index:idx_brand_product"`
	Brand       *Brand    `json:"brand" gorm:"foreignKey:BrandID"`
	ProductID   uint      `json:"product_id" gorm:"not null"`
	PriceListID uint      `json:"pricelist_id" gorm:"not null"`
	StartDate   time.Time `json:"start_date" gorm:"type:date; not null"`
	EndDate     time.Time `json:"end_date" gorm:"type:date; not null"`
	Priority    uint      `json:"priority" gorm:"not null"`
	Price       int       `json:"price" gorm:"not null"`
	Currency    string    `json:"currency" gorm:"not null"`
}
