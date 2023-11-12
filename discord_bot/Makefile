BOT_BINARY=bot
MIGRATION_BINARY=migrate

bot:
	go build -o bin/$(BOT_BINARY) ./cmd/bot.go
	cd bin; ./$(BOT_BINARY)

migration:
	go build -o bin/$(MIGRATION_BINARY) ./cmd/migrate.go
	cd bin; ./$(MIGRATION_BINARY)