BINARY_NAME=clashsync

run:
	go build -o bin/$(BINARY_NAME)_dev ./cmd/run.go
	cd bin; ./$(BINARY_NAME)_dev

build:
	go build -o bin/$(BINARY_NAME) ./cmd/run.go

clean:
	go clean
	rm -rf bin/$(BINARY_NAME)*