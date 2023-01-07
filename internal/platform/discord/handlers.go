package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/discord-bot/internal/service"
)

const PREFIX = "/"

func (b *Bot) handlerCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == b.botID || user.Bot {
		return
	}

	content := message.Content

	if len(content) <= len(PREFIX) {
		return
	}

	if content[:len(PREFIX)] != PREFIX {
		return
	}

	content = content[len(PREFIX):]

	if len(content) < 1 {
		return
	}

	args := strings.Fields(content)

	name := strings.ToLower(args[0])

	command, found := b.commands.GetCommand(name)
	if !found {
		return
	}

	channel, err := b.dg.State.Channel(message.ChannelID)
	if err != nil {
		b.logger.Errorf("err with getting channel: %v", err)
		return
	}

	guild, err := b.dg.State.Guild(channel.GuildID)
	if err != nil {
		b.logger.Errorf("err with getting guild: %v", err)
		return
	}

	ctx := service.NewContext(
		session,
		guild,
		channel,
		user,
		b.commands.GetCommands(),
		name,
		args[1:],
		b.logger,
	)

	c := *command

	c(*ctx)
}
