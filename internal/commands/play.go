package commands

import "github.com/discord-bot/internal/service"

func Play(ctx service.Context) {
	if _, err := ctx.Discord.ChannelMessageSend(ctx.Channel.ID, "play :)"); err != nil {
		return
	}
}
