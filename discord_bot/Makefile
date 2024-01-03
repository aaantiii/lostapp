BOT_BINARY=bot

bot:
	go build -o bin/$(BOT_BINARY)_dev ./cmd/bot.go
	cd bin; ./$(BOT_BINARY)_dev

build:
	go build -o bin/$(BOT_BINARY) ./cmd/bot.go
	GOARCH=arm64 go build -o bin/$(BOT_BINARY)_arm64 ./cmd/bot.go

clean:
	go clean
	rm -rf bin/$(BOT_BINARY)*