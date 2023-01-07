package commands

import "github.com/discord-bot/internal/service"

// TODO err logger
func Play(ctx service.Context) {
	vc := ctx.GetVoiceChannel()
	if vc == nil {
		ctx.Reply("You must be in a voice channel")
		return
	}

	voiceConnection := ctx.JoinVoiceChannel(ctx.Guild.ID, vc.ID)
	if voiceConnection == nil {
		ctx.Reply("An error occurred!")
		return
	}

	// go Start()

	ctx.Reply("play :)")
}
