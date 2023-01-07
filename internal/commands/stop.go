package commands

import (
	"github.com/discord-bot/internal/service"
)

func Stop(ctx service.Context) {
	vc := ctx.GetVoiceChannel()
	if vc == nil {
		ctx.Reply("You must be in a voice channel")
		return
	}

	ctx.Discord.RLock()
	voice, ok := ctx.Discord.VoiceConnections[vc.GuildID]
	ctx.Discord.RUnlock()
	if !ok {
		return
	}

	voice.Close()

	if err := voice.Disconnect(); err != nil {
		ctx.Reply("An error occurred!")
		return
	}
}
