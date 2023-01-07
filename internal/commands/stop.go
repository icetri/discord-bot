package commands

import (
	"github.com/discord-bot/internal/service"
)

// TODO err logger
func Stop(ctx service.Context) {
	// Проверка что уже существует в канале

	vc := ctx.GetVoiceChannel()
	if vc == nil {
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
		return
	}
}
