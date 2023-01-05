package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/discord-bot/internal/commands"
	"github.com/discord-bot/internal/config"
	"github.com/discord-bot/internal/platform/discord"
	"github.com/discord-bot/internal/service"
)

var (
	cmds *service.Commands
)

func Run() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
		return
	}

	cmds = service.NewCommands()
	registerCommands()

	dg, err := discord.NewBot(cfg, cmds)
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
	//	c.Register("help", HelpCommand, "TODO")
	cmds.Register("play", commands.Play, "TODO")
	//	c.Register("stop", StopCommand, "TODO")
}
