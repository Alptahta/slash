run:
	go run main.go

unit-tests:
	go test -v ./...

unit-test-coverage:
	go test -v ./... -coverprofile coverage.out
	go tool cover -html coverage.out -o coverage.html

build-image:
	docker build --tag slash .