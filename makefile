build:
	@go build -o bin/go-fuck-this-shit

run: build
	@./bin/go-fuck-this-shit

test: 
	@go test -v ./...

 
