install:
	go mod download

check:
	gofmt -e . > /dev/null