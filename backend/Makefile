BINARY_NAME=backend

serve:
	go build -o bin/$(BINARY_NAME)_dev ./cmd/serve.go
	cd bin; ./$(BINARY_NAME)_dev

build:
	go build -o bin/$(BINARY_NAME) ./cmd/serve.go

clean:
	go clean
	rm -rf bin/$(BINARY_NAME)*