build:
	go build -v ./src/cmd/...
tidy:
	go mod tidy
lint:
	golangci-lint run
test:
	go test -v ./src/...