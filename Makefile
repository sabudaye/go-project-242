build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

test:
	go test ./...

lint:
	$(go env GOPATH)/bin/golangci-lint run

lint-fix:
	$(go env GOPATH)/bin/golangci-lint run --fix
