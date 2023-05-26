package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/lelebus/inditext-go-backend/test/mocks"
	"github.com/stretchr/testify/require"

	"github.com/lelebus/inditext-go-backend/internal/domain"
)

////////////////////////
//
//  Queries
//
////////////////////////

func TestGetBrand(t *testing.T) {

	t.Run("should return 404 not found", func(t *testing.T) {
		nonExistentID := 12345
	
		url := fmt.Sprintf("http://localhost:%s/brand?id=%d", port, nonExistentID)
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
	
		// Check the response status code
		if res.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status %v, but got %v", http.StatusNotFound, res.StatusCode)
		}
	})	

	t.Run("should return brand struct", func(t *testing.T) {
		firstBrand := mocks.MockedBrands[0]

		url := fmt.Sprintf("http://localhost:%s/brand?id=%d", port, firstBrand.ID)
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		var data domain.Brand
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			panic(err)
		}

		// assert results
		require.Equal(t, *firstBrand, data)
	})

}
