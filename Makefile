build:
	@go build -o bin/portfolion

run: build
	@./bin/portfolion

test: 
	@go test -v ./...