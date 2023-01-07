package commands

import "github.com/discord-bot/internal/service"

// TODO err logger
func Play(ctx service.Context) {
	vc := ctx.GetVoiceChannel()
	if vc == nil {
		return
	}

	_, err := ctx.JoinVoiceChannel(ctx.Guild.ID, vc.ID)
	if err != nil {
		return
	}

	// go Start()

	ctx.Reply("play :)")
}
