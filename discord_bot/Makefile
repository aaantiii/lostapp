BOT_BINARY=lostmanager

bot:
	go build -o bin/$(BOT_BINARY)_dev ./cmd/bot.go
	cd bin; ./$(BOT_BINARY)_dev

build:
	go build -o bin/$(BOT_BINARY) ./cmd/bot.go

clean:
	go clean
	rm -rf bin/$(BOT_BINARY)*