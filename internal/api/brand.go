package api

import (
	"encoding/json"
	"net/http"

	"github.com/lelebus/inditext-go-backend/internal/app"
)

func GetBrand(w http.ResponseWriter, req *http.Request) {
	// Get the "id" query parameter value
	id := req.URL.Query().Get("id")

	brand, err := app.GetBrandByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if brand == nil {
		http.Error(w, "Brand not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(brand)
}
