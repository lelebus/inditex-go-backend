# Inditex Backend Test Case

The project is a REST API backend service in Golang.

## Usage

```bash
# To run the server at localhost:8080
make run

# To run the server at localhost:8080 but preload the database with mock data
make run_with_mocks

# To run the tests for the REST API endpoints
make it_api
```

Once the server has been started, you can play around with the following queries:

- http://localhost:8080/brand?id=1
- http://localhost:8080/pvp?brand_id=1&product_id=35455&application_date=2020-06-16-21.00.00

### Notes

The backend lacks APIs for populating the database, rendering the server startup futile without mock data.

Nonetheless, incorporating an endpoint for this purpose is simple due to the availability of the service on the database layer.

Furthermore, the futility of starting the server without mock data stems from the database's lack of persistence, given its in-memory nature. However, transitioning to a persistent database is effortless, thereby enhancing the server's utility.

## Tests

The implementation of the 5 specified tests can be found in the `test/api/price_test.go` file.

These tests are categorized as integration tests since they rely on the server being active and the database layer being properly configured. To accommodate this requirement, they have been organized within a dedicated package specifically designed for API testing.

## Further Notes:

The instruction indicate the following

```
Accept as input parameters: application date, product identifier, string identifier
```

However, the prompt does not provide a specific definition or explanation of what the `string identifier` refers to.

I used the following as inputs for the query to retrieve the retail price:

- brand_id
- product_id
- application_date
