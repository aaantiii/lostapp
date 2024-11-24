BOT_BINARY=lostmanager

bot:
	go build -o bin/$(BOT_BINARY)_dev ./cmd/bot.go
	cd bin; ./$(BOT_BINARY)_dev

build-linux:
	go build -o bin/$(BOT_BINARY) ./cmd/bot.go

build-win:
	go build -o bin/$(BOT_BINARY).exe ./cmd/bot.go

bot-win:
	go build -o bin/$(BOT_BINARY).exe ./cmd/bot.go
	cd bin && $(BOT_BINARY).exe

clean:
	go clean
	rm -rf bin/$(BOT_BINARY)*