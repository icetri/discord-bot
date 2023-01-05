package discord

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"

	"github.com/discord-bot/internal/config"
	"github.com/discord-bot/internal/service"
)

type Bot struct {
	dg       *discordgo.Session
	commands *service.Commands
	botID    string
	ownerID  string
}

func NewBot(cfg *config.Config, commands *service.Commands) (*Bot, error) {
	discord, err := discordgo.New(cfg.BotToken)
	if err != nil {
		return nil, fmt.Errorf("error creating discord session: %w", err)
	}

	return &Bot{
		dg:       discord,
		ownerID:  cfg.OwnerID,
		commands: commands,
	}, nil
}

func (b *Bot) ping() error {
	user, err := b.dg.User(b.ownerID)
	if err != nil {
		return fmt.Errorf("error with take user details: %w", err)
	}

	log.Printf("Authorization user: %s", user.Username)
	b.botID = user.ID

	return nil
}

func (b *Bot) Start() error {
	err := b.dg.Open()
	if err != nil {
		return fmt.Errorf("error opening connection: %w", err)
	}

	if err := b.ping(); err != nil {
		return fmt.Errorf("error with ping: %w", err)
	}

	b.handlerUpdate()

	return nil
}

func (b *Bot) Close() error {
	if err := b.dg.Close(); err != nil {
		return fmt.Errorf("error close connection: %w", err)
	}

	return nil
}

func (b *Bot) handlerUpdate() {
	b.dg.AddHandler(b.handlerCommand)
}
