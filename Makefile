build:
	@go build -o bin/paperspls ./cmd/pp/

run: build
	@./bin/paperspls

test: 
	@go test -v ./...

