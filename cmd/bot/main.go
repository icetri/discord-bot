package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/discord-bot/internal/service"
	"github.com/discord-bot/pkg/config"
	"github.com/discord-bot/pkg/discord"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
		return
	}

	commands := service.NewCommands()

	dg, err := discord.NewBot(cfg, commands)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := dg.Start(); err != nil {
		log.Fatal(err)
		return
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	if err := dg.Close(); err != nil {
		log.Fatal(err)
	}
}
