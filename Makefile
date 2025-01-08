build:
	@go build -o bin/paperspls ./cmd/

run: build
	@./bin/paperspls

test: 
	@go test -v ./...

