package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/lelebus/inditext-go-backend/internal/app"
)

func GetPrice(w http.ResponseWriter, req *http.Request) {
	// Parse query parameters
	brandId := req.URL.Query().Get("brand_id")
	productId := req.URL.Query().Get("product_id")
	applicationDate, err := time.Parse("2006-01-02-15.04.05", req.URL.Query().Get("application_date"))
	if err != nil {
		http.Error(w, "Invalid application date", http.StatusBadRequest)
		return
	}

	// Get the price
	price, err := app.GetPriceByApplicationDate(brandId, productId, applicationDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if price == nil {
		http.Error(w, "No price found for this product", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(price)
}
