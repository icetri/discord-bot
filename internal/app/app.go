package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/discord-bot/internal/commands"
	"github.com/discord-bot/internal/config"
	"github.com/discord-bot/internal/platform/discord"
	"github.com/discord-bot/internal/service"
	"github.com/discord-bot/pkg/logger"
)

var (
	cmds *service.Commands
)

func Run() {
	log := logger.New()

	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
		return
	}

	cmds = service.NewCommands()
	registerCommands()

	dg, err := discord.NewBot(cfg, cmds, log)
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

func registerCommands() {
	cmds.Register("help", commands.Help, "TODO")
	cmds.Register("play", commands.Play, "TODO")
	cmds.Register("stop", commands.Stop, "TODO")
}
