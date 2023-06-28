run:
	go run main.go

unit-tests:
	go test -v ./...

unit-test-coverage:
	go test -v ./... -coverprofile coverage.out
	go tool cover -html coverage.out -o coverage.html

dockerized:
	# docker build --tag slash .
	# docker build -t slash:v1.0 .
	docker build -t slash:multistage -f Dockerfile.multistage .
	docker run -d -p 8080:8080 --name slashCont slash:multistage


# To purge your computer from this project's related docker
purge:
	docker stop slashCont
	docker rm slashCont
	docker image rm slash:latest
	docker image rm slash:multistage