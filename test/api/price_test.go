package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/lelebus/inditext-go-backend/internal/domain"
	"github.com/lelebus/inditext-go-backend/test/mocks"
	"github.com/stretchr/testify/require"
)

////////////////////////
//
//  Queries
//
////////////////////////

func TestGetPrice(t *testing.T) {

	t.Run("test 1", func(t *testing.T) {
		tmstp := "2020-06-14-10.00.00"
		brandId := 1
		productId := 35455
		url := fmt.Sprintf("http://localhost:%s/pvp?brand_id=%d&product_id=%d&application_date=%s", port, brandId, productId, tmstp)
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Check the response status code
		if res.StatusCode != http.StatusOK {
			t.Fatalf("Expected status %v, but got %v", http.StatusOK, res.StatusCode)
		}

		var data domain.Price
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			panic(err)
		}

		// assert results
		expectedResult := mocks.MockedPrices[0]
		require.Equal(t, *expectedResult, data)

		// reinforce the test, in case the mocked data changes
		expectedPrice := 3550
		require.Equal(t, expectedPrice, data.Price)
	})

	t.Run("test 2", func(t *testing.T) {
		tmstp := "2020-06-14-16.00.00"
		brandId := 1
		productId := 35455
		url := fmt.Sprintf("http://localhost:%s/pvp?brand_id=%d&product_id=%d&application_date=%s", port, brandId, productId, tmstp)
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Check the response status code
		if res.StatusCode != http.StatusOK {
			t.Fatalf("Expected status %v, but got %v", http.StatusOK, res.StatusCode)
		}

		var data domain.Price
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			panic(err)
		}

		// assert results
		expectedResult := mocks.MockedPrices[1]
		require.Equal(t, *expectedResult, data)

		// reinforce the test, in case the mocked data changes
		expectedPrice := 2545
		require.Equal(t, expectedPrice, data.Price)
	})

	t.Run("test 3", func(t *testing.T) {
		tmstp := "2020-06-14-21.00.00"
		brandId := 1
		productId := 35455
		url := fmt.Sprintf("http://localhost:%s/pvp?brand_id=%d&product_id=%d&application_date=%s", port, brandId, productId, tmstp)
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Check the response status code
		if res.StatusCode != http.StatusOK {
			t.Fatalf("Expected status %v, but got %v", http.StatusOK, res.StatusCode)
		}

		var data domain.Price
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			panic(err)
		}

		// assert results
		expectedResult := mocks.MockedPrices[0]
		require.Equal(t, *expectedResult, data)

		// reinforce the test, in case the mocked data changes
		expectedPrice := 3550
		require.Equal(t, expectedPrice, data.Price)
	})

	t.Run("test 4", func(t *testing.T) {
		tmstp := "2020-06-15-10.00.00"
		brandId := 1
		productId := 35455
		url := fmt.Sprintf("http://localhost:%s/pvp?brand_id=%d&product_id=%d&application_date=%s", port, brandId, productId, tmstp)
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Check the response status code
		if res.StatusCode != http.StatusOK {
			t.Fatalf("Expected status %v, but got %v", http.StatusOK, res.StatusCode)
		}

		var data domain.Price
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			panic(err)
		}

		// assert results
		expectedResult := mocks.MockedPrices[2]
		require.Equal(t, *expectedResult, data)

		// reinforce the test, in case the mocked data changes
		expectedPrice := 3050
		require.Equal(t, expectedPrice, data.Price)
	})

	t.Run("test 5", func(t *testing.T) {
		tmstp := "2020-06-16-21.00.00"
		brandId := 1
		productId := 35455
		url := fmt.Sprintf("http://localhost:%s/pvp?brand_id=%d&product_id=%d&application_date=%s", port, brandId, productId, tmstp)
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Check the response status code
		if res.StatusCode != http.StatusOK {
			t.Fatalf("Expected status %v, but got %v", http.StatusOK, res.StatusCode)
		}

		var data domain.Price
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			panic(err)
		}

		// assert results
		expectedResult := mocks.MockedPrices[3]
		require.Equal(t, *expectedResult, data)

		// reinforce the test, in case the mocked data changes
		expectedPrice := 3895
		require.Equal(t, expectedPrice, data.Price)
	})

}
