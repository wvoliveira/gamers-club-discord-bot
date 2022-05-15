.DEFAULT_GOAL := build

.PHONY: build
build:
	go build -o gcbot cmd/bot/main.go

.PHONY: raspberry
raspberry:
	GOOS=linux GOARCH=arm GOARM=5 go build -o gcbot cmd/bot/main.go
