.PHONY: install check

install:
	go mod download

check:
	go vet ./...
	go build -v -o /dev/null ./...