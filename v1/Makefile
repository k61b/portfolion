build:
	@go build -o bin/portfolion

run: build
	@./bin/portfolion

test: 
	@go test -v ./...

coverage:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out