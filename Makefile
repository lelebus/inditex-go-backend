run:
	go run cmd/inditex/main.go

run_with_mocks:
	MOCKS=true go run cmd/inditex/main.go
	
it_api:
	go test ./test/api/...