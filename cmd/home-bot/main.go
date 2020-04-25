package main

import (
	"log"
	"os"

	"github.com/szeist/home-bot/internal/config"
	"github.com/szeist/home-bot/internal/keybasebot"
	"github.com/szeist/home-bot/internal/messagehandler"
)

func main() {
	logger := log.New(os.Stderr, "home-bot:main", log.LstdFlags)

	cfg := config.NewFromEnv()
	msgHandler := messagehandler.New(cfg)
	app := keybasebot.New(cfg.Keybase)

	err := app.Start()
	if err != nil {
		logger.Fatalf("Keybase bot creation error: %s", err)
	}

	go app.HandleMessages(msgHandler.Message)
	go app.HandleResponses(msgHandler.Response)

	msgHandler.Start()
}
